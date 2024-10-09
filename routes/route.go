package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/pejuangreceh/test-gemini.git/config"
	"github.com/pejuangreceh/test-gemini.git/controllers"
)

func Configure(s *config.Server) *echo.Echo {
	e := echo.New()
	v1 := "/api/elloy/v1"
	baseController := controllers.NewBaseController(s)
	e.POST(v1+"/test", baseController.TextGenerate)
	e.POST(v1+"/upload", baseController.TextGenerateFromImage)
	e.GET(v1+"/check-env", func(ctx echo.Context) error {
		fmt.Println("test ", s)
		return nil
	})

	return e
}
