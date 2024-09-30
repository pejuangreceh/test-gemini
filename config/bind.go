package config

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type request interface {
	Validate(validationType string) interface{}
}

func (s *Server) Bind(c echo.Context, i request) error {
	if i == nil {
		return nil
	}

	binder := new(echo.DefaultBinder)
	if err := binder.BindBody(c, i); err != nil {
		logrus.Error(c, logrus.Fields{
			"tag":   "Server.Bind.03",
			"error": err.Error(),
		}, "default bind body failed")

		return echo.NewHTTPError(http.StatusBadRequest, "ERROR_BINDING")
	}

	if err := binder.BindQueryParams(c, i); err != nil {
		logrus.Error(c, logrus.Fields{
			"tag":   "Server.Bind.04",
			"error": err.Error(),
		}, "default bind query param failed")

		return echo.NewHTTPError(http.StatusBadRequest, "ERROR_BINDING")
	}

	if err := binder.BindPathParams(c, i); err != nil {
		logrus.Error(c, logrus.Fields{
			"tag":   "Server.Bind.05",
			"error": err.Error(),
		}, "default bind path param failed")

		return echo.NewHTTPError(http.StatusBadRequest, "ERROR_BINDING")
	}

	if err := binder.BindHeaders(c, i); err != nil {
		logrus.Error(c, logrus.Fields{
			"tag":   "Server.Bind.06",
			"error": err.Error(),
		}, "default bind header failed")

		return echo.NewHTTPError(http.StatusBadRequest, "ERROR_BINDING")
	}

	// if i.Validate(s.Config.Validation) == nil {
	// 	return nil
	// }

	// if s.Config.Validation == "thedevsaddam" {
	// 	var e url.Values

	// 	opts := govalidator.Options{
	// 		Data:  i,
	// 		Rules: i.Validate(s.Config.Validation).(govalidator.MapData),
	// 	}

	// 	v := govalidator.New(opts)

	// 	e = v.ValidateStruct()
	// 	if len(e) != 0 {
	// 		if e.Get("_error") != "" {
	// 			logrus.Error(c, logrus.Fields{
	// 				"tag":   "Server.Bind.07",
	// 				"error": e.Get("_error"),
	// 			}, "get message failed")

	// 			return echo.NewHTTPError(http.StatusBadRequest, e.Get("_error"))
	// 		}

	// 		logrus.Error(c, logrus.Fields{
	// 			"tag":   "Server.Bind.08",
	// 			"error": e,
	// 		}, "validator failed")

	// 		return echo.NewHTTPError(http.StatusBadRequest, e)
	// 	}
	// } else {
	// 	e := i.Validate(s.Config.Validation)

	// 	return echo.NewHTTPError(http.StatusBadRequest, e)
	// }

	return nil
}
