package main

import (
	"github.com/riny/demo-go-gin/app/router"
	"log/slog"
)

func main() {
	engine := router.New()

	err := engine.Run()
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
