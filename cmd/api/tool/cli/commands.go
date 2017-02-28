// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/markusklems/p7/cmd/api/design
// --out=$(GOPATH)/src/github.com/markusklems/p7/cmd/api
// --version=v1.1.0-dirty
//
// API "p7": CLI Commands
//
// The content of this file is auto-generated, DO NOT MODIFY

package cli

import (
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/markusklems/p7/cmd/api/client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type (
	// HealthHealthCommand is the command line data structure for the health action of health
	HealthHealthCommand struct {
		PrettyPrint bool
	}

	// CodeLambdaCommand is the command line data structure for the code action of lambda
	CodeLambdaCommand struct {
		// Lambda ID
		LambdaID    int
		PrettyPrint bool
	}

	// CreateLambdaCommand is the command line data structure for the create action of lambda
	CreateLambdaCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// DeleteLambdaCommand is the command line data structure for the delete action of lambda
	DeleteLambdaCommand struct {
		LambdaID    int
		PrettyPrint bool
	}

	// ListLambdaCommand is the command line data structure for the list action of lambda
	ListLambdaCommand struct {
		PrettyPrint bool
	}

	// ShowLambdaCommand is the command line data structure for the show action of lambda
	ShowLambdaCommand struct {
		// Lambda ID
		LambdaID    int
		PrettyPrint bool
	}

	// UpdateLambdaCommand is the command line data structure for the update action of lambda
	UpdateLambdaCommand struct {
		Payload     string
		ContentType string
		LambdaID    int
		PrettyPrint bool
	}

	// WatchLambdaCommand is the command line data structure for the watch action of lambda
	WatchLambdaCommand struct {
		LambdaID    int
		PrettyPrint bool
	}

	// DownloadCommand is the command line data structure for the download command.
	DownloadCommand struct {
		// OutFile is the path to the download output file.
		OutFile string
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "code",
		Short: `Retrieve lambda code with given id`,
	}
	tmp1 := new(CodeLambdaCommand)
	sub = &cobra.Command{
		Use:   `lambda ["/p7/lambdas/LAMBDA_ID/actions/code"]`,
		Short: `A microservice function named lambda`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "create",
		Short: `Creates a lambda`,
	}
	tmp2 := new(CreateLambdaCommand)
	sub = &cobra.Command{
		Use:   `lambda ["/p7/lambdas"]`,
		Short: `A microservice function named lambda`,
		Long: `A microservice function named lambda

Payload example:

{
   "code": "f5bzyr60yi",
   "name": "rs5k"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "delete",
		Short: ``,
	}
	tmp3 := new(DeleteLambdaCommand)
	sub = &cobra.Command{
		Use:   `lambda ["/p7/lambdas/LAMBDA_ID"]`,
		Short: `A microservice function named lambda`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "health",
		Short: `Perform health check.`,
	}
	tmp4 := new(HealthHealthCommand)
	sub = &cobra.Command{
		Use:   `health ["/p7/_ah/health"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `List all avilable lambdas`,
	}
	tmp5 := new(ListLambdaCommand)
	sub = &cobra.Command{
		Use:   `lambda ["/p7/lambdas"]`,
		Short: `A microservice function named lambda`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `Retrieve lambda with given id`,
	}
	tmp6 := new(ShowLambdaCommand)
	sub = &cobra.Command{
		Use:   `lambda ["/p7/lambdas/LAMBDA_ID"]`,
		Short: `A microservice function named lambda`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "update",
		Short: ``,
	}
	tmp7 := new(UpdateLambdaCommand)
	sub = &cobra.Command{
		Use:   `lambda ["/p7/lambdas/LAMBDA_ID"]`,
		Short: `A microservice function named lambda`,
		Long: `A microservice function named lambda

Payload example:

{
   "code": "f5bzyr60yi",
   "name": "rs5k"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "watch",
		Short: `Retrieve lambda with given id`,
	}
	tmp8 := new(WatchLambdaCommand)
	sub = &cobra.Command{
		Use:   `lambda ["/p7/lambdas/LAMBDA_ID/watch"]`,
		Short: `A microservice function named lambda`,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)

	dl := new(DownloadCommand)
	dlc := &cobra.Command{
		Use:   "download [PATH]",
		Short: "Download file with given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dl.Run(c, args)
		},
	}
	dlc.Flags().StringVar(&dl.OutFile, "out", "", "Output file")
	app.AddCommand(dlc)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if rpath == "/" {
		fnf = c.Download
		if outfile == "" {
			outfile = "index.html"
		}
		goto found
	}
	if rpath == "/swagger.json" {
		fnf = c.DownloadSwaggerJSON
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	if strings.HasPrefix(rpath, "/css/") {
		fnd = c.DownloadCSS
		rpath = rpath[5:]
		if outfile == "" {
			_, outfile = path.Split(rpath)
		}
		goto found
	}
	if strings.HasPrefix(rpath, "/img/") {
		fnd = c.DownloadImg
		rpath = rpath[5:]
		if outfile == "" {
			_, outfile = path.Split(rpath)
		}
		goto found
	}
	if strings.HasPrefix(rpath, "/js/") {
		fnd = c.DownloadJs
		rpath = rpath[4:]
		if outfile == "" {
			_, outfile = path.Split(rpath)
		}
		goto found
	}
	if strings.HasPrefix(rpath, "/swagger-ui/") {
		fnd = c.DownloadSwaggerUI
		rpath = rpath[12:]
		if outfile == "" {
			_, outfile = path.Split(rpath)
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

// Run makes the HTTP request corresponding to the HealthHealthCommand command.
func (cmd *HealthHealthCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/p7/_ah/health"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.HealthHealth(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *HealthHealthCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the CodeLambdaCommand command.
func (cmd *CodeLambdaCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/p7/lambdas/%v/actions/code", cmd.LambdaID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CodeLambda(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CodeLambdaCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var lambdaID int
	cc.Flags().IntVar(&cmd.LambdaID, "lambda_id", lambdaID, `Lambda ID`)
}

// Run makes the HTTP request corresponding to the CreateLambdaCommand command.
func (cmd *CreateLambdaCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/p7/lambdas"
	}
	var payload client.LambdaPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateLambda(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateLambdaCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the DeleteLambdaCommand command.
func (cmd *DeleteLambdaCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/p7/lambdas/%v", cmd.LambdaID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.DeleteLambda(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *DeleteLambdaCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var lambdaID int
	cc.Flags().IntVar(&cmd.LambdaID, "lambda_id", lambdaID, ``)
}

// Run makes the HTTP request corresponding to the ListLambdaCommand command.
func (cmd *ListLambdaCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/p7/lambdas"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListLambda(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListLambdaCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ShowLambdaCommand command.
func (cmd *ShowLambdaCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/p7/lambdas/%v", cmd.LambdaID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowLambda(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowLambdaCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var lambdaID int
	cc.Flags().IntVar(&cmd.LambdaID, "lambda_id", lambdaID, `Lambda ID`)
}

// Run makes the HTTP request corresponding to the UpdateLambdaCommand command.
func (cmd *UpdateLambdaCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/p7/lambdas/%v", cmd.LambdaID)
	}
	var payload client.LambdaPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.UpdateLambda(ctx, path, &payload, cmd.ContentType)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *UpdateLambdaCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
	var lambdaID int
	cc.Flags().IntVar(&cmd.LambdaID, "lambda_id", lambdaID, ``)
}

// Run establishes a websocket connection for the WatchLambdaCommand command.
func (cmd *WatchLambdaCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/p7/lambdas/%v/watch", cmd.LambdaID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	ws, err := c.WatchLambda(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}
	go goaclient.WSWrite(ws)
	goaclient.WSRead(ws)

	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *WatchLambdaCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var lambdaID int
	cc.Flags().IntVar(&cmd.LambdaID, "lambda_id", lambdaID, ``)
}
