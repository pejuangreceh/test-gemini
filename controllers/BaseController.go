package controllers

import (
	"io"
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
func (base BaseController) TextGenerateFromImage(c echo.Context) error {
	var (
		request requests.FileRequest
		tag     = "BaseController.TextGenerateFromImage."
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

	// Source
	file, err := c.FormFile("image_upload")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	if _, err := os.Stat("temp/img"); os.IsNotExist(err) {
		err = os.Mkdir("temp/img", 0664)
		if err != nil {
			logrus.Error(c, logrus.Fields{
				"tag":     tag + "02",
				"error":   err,
				"request": nil,
				"data":    nil,
			}, "INTERNAL_SERVER_ERROR")
		}
	}
	// defer os.RemoveAll("temp/img")
	basePath := "temp/img/"
	dst, err := os.Create(basePath + file.Filename)
	if err != nil {
		return err
	}
	// defer os.RemoveAll("temp/img")
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	// responses.ResponseData
	return c.JSON(http.StatusOK, responses.ResponseData{
		Code:        strconv.Itoa(http.StatusOK),
		Description: "Success",
		Data:        nil,
	})
}
