package backend

import (
	"fmt"
	"github.com/kuno989/friday/backend/pkg"
	"github.com/kuno989/friday/backend/schema"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func (s *Server) UploadFile(c echo.Context) error {
	responseFile, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, schema.FileResponse{
			Message:     "파일이 없습니다",
			Description: "파일이 제대로 업로드 되었는지 확인해주세요",
		})
	}
	if responseFile.Size > s.Config.MaxFileSize  {
		return c.JSON(http.StatusRequestEntityTooLarge, schema.FileResponse{
			Message:     "업로드 실패",
			Description: "최대 파일 업로드 가능 용량 64mb",
			FileName:    responseFile.Filename,
			FileSize:    responseFile.Size,
		})
	}
	file, err := responseFile.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, schema.FileResponse{
			Message:     "Internal error",
			Description: "Internal error",
			FileName:    responseFile.Filename,
			FileSize:    responseFile.Size,
		})
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, schema.FileResponse{
			Message:     "파일 읽기 실패",
			Description: "Internal error",
			FileName:    responseFile.Filename,
			FileSize:    responseFile.Size,
		})
	}
	sha256 := pkg.NewSHA256(content)
	fmt.Println(sha256)
	return c.JSON(http.StatusBadRequest, "Filters not allowed") // 수정예정
}

func (s *Server)FileGetHandler(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, "Filters not allowed")
}