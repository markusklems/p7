//go:generate goagen bootstrap -d github.com/markusklems/p7/cmd/api/design

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/inconshreveable/log15"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
	"github.com/markusklems/p7/cmd/api/app"
	"github.com/markusklems/p7/cmd/api/controllers"
	"github.com/markusklems/p7/cmd/api/models"
)

var db *gorm.DB
var logger log15.Logger
var ldb *models.LambdaDB

func main() {
	var addr = flag.String("addr", "0.0.0.0:8888", "The ip / host and port of the application. (string)")
	//var inside = flag.Bool("inside", true, "Binary runs inside a Kubernetes cluster. (bool)")
	flag.Parse() // parse the flags
	password := getenv("SECRET_PASSWORD", "p7password")
	host := getenv("P7_DB01_MYSQL_SERVICE_HOST", "127.0.0.1")
	port, err := strconv.Atoi(getenv("P7_DB01_MYSQL_SERVICE_PORT", "3306"))
	if err != nil {
		panic(err)
	}

	// Create db context
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&autocommit=true", "user", password, host, port, "p7")
	fmt.Println(url)
	db, err = gorm.Open("mysql", url)
	if err != nil {
		fmt.Printf("Couldn't connect to database, proceed without database. Error: %s\n", err)
	} else {
		db.LogMode(true)

		db.DropTable(&models.Lambda{})
		db.AutoMigrate(&models.Lambda{})
		db.DB().SetMaxOpenConns(50)

		ldb = models.NewLambdaDB(db)
	}

	// Create service
	service := goa.New("p7")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "lambda" controller
	fc := controllers.NewLambda(service, ldb)
	app.MountLambdaController(service, fc)

	// Mount "health" controller
	hc := controllers.NewHealth(service)
	app.MountHealthController(service, hc)

	// Mount "js" controller
	jc := controllers.NewJs(service)
	app.MountJsController(service, jc)

	// Mount "public" controller
	pc := controllers.NewPublic(service)
	app.MountPublicController(service, pc)

	// Mount "img" controller
	ic := controllers.NewImg(service)
	app.MountImgController(service, ic)

	// Mount "img" controller
	cc := controllers.NewCSS(service)
	app.MountCSSController(service, cc)

	// Mount "swagger" controller
	sc := controllers.NewSwagger(service)
	app.MountSwaggerController(service, sc)

	// Start service
	if err := service.ListenAndServe(*addr); err != nil {
		service.LogError("startup", "err", err)
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
