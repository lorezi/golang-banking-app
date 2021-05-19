package main

import (
	"github.com/lorezi/golang-bank-app/app"
	"github.com/lorezi/golang-bank-app/logger"
)

func main() {

	logger.Info("Starting the application... 🤝")
	app.Start()
}
