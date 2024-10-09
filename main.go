package main

import (
	"github.com/pejuangreceh/test-gemini.git/config"
	"github.com/pejuangreceh/test-gemini.git/routes"
	"github.com/sirupsen/logrus"
	// application "test-gemini"
)

func main() {

	cfg := config.New()
	app := config.NewServer(cfg)
	e := routes.Configure(app)
	err := app.Start(e)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"tag":   "App.Start.01",
			"error": err,
		}).Error("start server failed")
	}

}
