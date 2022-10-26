package main

import (
	"github.com/y-sugiyama654/banking-lib/logger"
	"github.com/y-sugiyama654/banking/app"
)

func main() {
	logger.Info("Starting application.")
	app.Start()
}
