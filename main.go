package main

import (
	"github.com/y-sugiyama654/banking/app"
	"github.com/y-sugiyama654/banking/logger"
)

func main() {
	logger.Info("Starting application.")
	app.Start()
}
