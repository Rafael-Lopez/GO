package main

import (
	"github.com/Rafael-Lopez/GO/RESTAPI/banking/app"
	"github.com/Rafael-Lopez/GO/RESTAPI/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
