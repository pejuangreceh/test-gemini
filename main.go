package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pejuangreceh/test-gemini.git/config"
	"github.com/pejuangreceh/test-gemini.git/controllers"
	// application "test-gemini"
)

func main() {

	e := echo.New()
	v1 := "/api/elloy/v1"
	var s *config.Server
	baseController := controllers.NewBaseController(s)
	e.POST(v1+"test", baseController.TextGenerate)
	e.Logger.Fatal(e.Start("0.0.0.0:" + "8001"))

}
