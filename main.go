package main

import (
	"github.com/riny/demo-go-gin/app/router"
	"log/slog"
)

// @title TodoList API Documentation
// @version v1.0
// @description TodoList API Documentation
// @termsOfService https://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://127.0.0.1:8080
// @BasePath /

func main() {
	engine := router.New()

	err := engine.Run()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
