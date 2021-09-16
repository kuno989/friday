package files

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func UploadFile(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, "Filters not allowed") // 수정예정
}

func FileGetHandler(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, "Filters not allowed")
}