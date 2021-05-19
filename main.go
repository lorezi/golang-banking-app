package main

import (
	"github.com/lorezi/golang-bank-app/app"
	"github.com/lorezi/golang-bank-app/logger"
)

func main() {

	// log.Println("Starting our application...🧜🏽🧜🏽🤗")
	logger.Log.Info("Starting the application... 🤝")
	app.Start()
}
