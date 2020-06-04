package main

import (
	"github.com/OoThan/go-api-test/api/config"

	"github.com/OoThan/go-api-test/api/app"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
