package response

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"practice/golang/internal/logger"
)

var log = logger.Logger()

type Response struct {
	Status       bool        `json:"status"`
	Data         interface{} `json:"data,omitempty"`
	ErrorCode    string      `json:"error_code,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
}

func Error(c echo.Context, err error) error {
	log.Error(err)
	return c.JSON(http.StatusInternalServerError,"error")
}

func Success(c echo.Context, data interface{}) error {
	var r = Response{
		Status:      true,
		Data:        data,
	}
	return c.JSON(http.StatusOK, &r)
}
