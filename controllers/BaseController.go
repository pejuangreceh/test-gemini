package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pejuangreceh/test-gemini.git/config"
	"github.com/pejuangreceh/test-gemini.git/gemini"
	"github.com/pejuangreceh/test-gemini.git/requests"
	"github.com/pejuangreceh/test-gemini.git/responses"
	"github.com/sirupsen/logrus"
)

type BaseController struct {
	*config.Server
}

func NewBaseController(s *config.Server) BaseController {
	return BaseController{s}
}

func (base BaseController) TextGenerate(c echo.Context) error {
	var (
		request requests.TextRequest
		tag     = "BaseController.TextGenerate."
	)
	if err := base.Bind(c, &request); err != nil {
		log.Fatal(logrus.Fields{
			"tag":     tag + "01",
			"error":   err,
			"request": request,
			"data":    nil,
		}, "BAD_REQUEST")
		return c.JSON(http.StatusBadRequest, responses.ResponseData{
			Code:        strconv.Itoa(http.StatusBadRequest),
			Description: "Bad Request coy",
			Data:        err,
		})
	}
	fmt.Println("masuk === ", os.Getenv("API_KEY"))

	jsonData, err := gemini.NewClient(request.Question)
	if err != nil {
		log.Fatal(logrus.Fields{
			"tag":     tag + "02",
			"error":   err,
			"request": request,
			"data":    nil,
		}, "INTERNAL_SERVER_ERROR")
		return c.JSON(http.StatusInternalServerError, responses.ResponseData{
			Code:        strconv.Itoa(http.StatusInternalServerError),
			Description: "Internal Server Error xixixixi",
			Data:        nil,
		})
	}
	// responses.ResponseData
	return c.JSON(http.StatusOK, responses.ResponseData{
		Code:        strconv.Itoa(http.StatusOK),
		Description: "Success",
		Data:        jsonData,
	})

}
