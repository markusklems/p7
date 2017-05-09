package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/markusklems/p7/cmd/api/app"
	"github.com/markusklems/p7/cmd/api/models"
	"golang.org/x/net/websocket"
)

// ErrDatabaseError is the error returned when a db query fails.
var ErrDatabaseError = goa.NewErrorClass("db_error", 500)

// LambdaController implements the lambda resource.
type LambdaController struct {
	*goa.Controller
	ldb       *models.LambdaDB
	netClient *http.Client
	logger    goa.LogAdapter
}

// NewLambda creates a lambda controller.
func NewLambda(service *goa.Service, ldb *models.LambdaDB, netClient *http.Client, logger goa.LogAdapter) *LambdaController {
	return &LambdaController{
		Controller: service.NewController("LambdaController"),
		ldb:        ldb,
		netClient:  netClient,
		logger:     logger,
	}
}

// Code runs the code action.
func (lc *LambdaController) Code(ctx *app.CodeLambdaContext) error {
	lambda, err := lc.ldb.Get(ctx.Context, ctx.LambdaID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.OKCode(lambda.LambdaToLambdaCode())
}

// Create runs the create action.
func (lc *LambdaController) Create(ctx *app.CreateLambdaContext) error {
	l := models.Lambda{}
	l.Name = ctx.Payload.Name
	l.Code = ctx.Payload.Code
	err := lc.ldb.Add(ctx.Context, &l)
	if err != nil {
		return ErrDatabaseError(err)
	}
	ctx.ResponseData.Header().Set("Location", app.LambdaHref(l.ID))

	//c := client.New(goaclient.HTTPClientDoer(http.DefaultClient))
	//c.UserAgent = "image-cli/0"
	//test := &client.CreateImagePayload{}
	//c.CreateImage(ctx.Context, "image/images/"+strconv.Itoa(l.ID), test, "application/json")

	return ctx.Created()
}

// Delete runs the delete action.
func (lc *LambdaController) Delete(ctx *app.DeleteLambdaContext) error {
	err := lc.ldb.Delete(ctx.Context, ctx.LambdaID)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()
}

// List runs the list action.
func (lc *LambdaController) List(ctx *app.ListLambdaContext) error {
	lambdas := lc.ldb.ListLambda(ctx.Context)
	return ctx.OK(lambdas)
}

// Run runs the run action.
func (lc *LambdaController) Run(ctx *app.RunLambdaContext) error {
	// LambdaController_Run: start_implement
	lambda, err := lc.ldb.OneLambda(ctx.Context, ctx.LambdaID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}
	lc.logger.Info("Run action with method: %s\n", ctx.Request.Method)
	if lambda.Method == ctx.Request.Method {
		resp, err := lc.netClient.PostForm("http://p7-kube:", url.Values{"q": {"github"}})
		if err != nil {
			lc.logger.Error("Couldn't reach Kube component: %q", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			lc.logger.Error("Couldn't decode Kube response: %q", err)
		}
		lc.logger.Info("Created new Kubernetes job: %s\n", string(body))
		return nil
	}
	return ctx.NotFound()
	// LambdaController_Run: end_implement
}

// Show runs the show action.
func (lc *LambdaController) Show(ctx *app.ShowLambdaContext) error {
	lambda, err := lc.ldb.OneLambda(ctx.Context, ctx.LambdaID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}
	lambda.Href = app.LambdaHref(lambda.ID)
	return ctx.OK(lambda)
}

// Update runs the update action.
func (lc *LambdaController) Update(ctx *app.UpdateLambdaContext) error {
	lambda, err := lc.ldb.Get(ctx.Context, ctx.LambdaID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	}
	payload := ctx.Payload
	if payload.Name != "" {
		lambda.Name = payload.Name
	}
	if payload.Code != "" {
		lambda.Code = payload.Code
	}
	err = lc.ldb.Update(ctx, lambda)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()
}

// Watch runs the watch action.
func (lc *LambdaController) Watch(ctx *app.WatchLambdaContext) error {
	lc.WatchWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// WatchWSHandler establishes a websocket connection to run the watch action.
func (lc *LambdaController) WatchWSHandler(ctx *app.WatchLambdaContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		watched := fmt.Sprintf("Lambda: %d", ctx.LambdaID)
		ws.Write([]byte(watched))
		io.Copy(ws, ws)
	}
}
