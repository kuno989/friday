package backend

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kuno989/friday/backend/pkg"
	"github.com/kuno989/friday/backend/schema"
	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	regex = "^[a-f0-9]{64}$"
)

func (s *Server) FileGetHandler(c echo.Context) error {
	sha256 := c.Param("sha256")
	sha256 = strings.ToLower(sha256)
	matched, _ := regexp.MatchString(regex, sha256)
	if !matched {
		return c.JSON(http.StatusBadRequest, schema.FileResponse{
			Message:     "sha256 포멧이 아닙니다",
			Description: "잘못된 hash 요청입니다 : " + sha256,
		})
	}
	file, err := s.ms.FileSearch(c.Request().Context(), sha256)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(http.StatusOK, schema.FileResponse{
				Sha256:      sha256,
				Message:     err.Error(),
				Description: "파일을 찾을 수 없습니다",
			})
		}
	}
	return c.JSON(http.StatusOK, file)
}

func (s *Server) UpdateFile(c echo.Context) error {
	sha256 := strings.ToLower(c.Param("sha256"))
	matched, _ := regexp.MatchString(regex, sha256)
	if !matched {
		return c.JSON(http.StatusBadRequest, schema.FileResponse{
			Message:     "sha256 포멧이 아닙니다",
			Description: "잘못된 hash 요청입니다 : " + sha256,
		})
	}
	b, err := ioutil.ReadAll(c.Request().Body)
	ctx := context.Background()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	file, err := s.ms.FileSearch(ctx, sha256)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(http.StatusOK, schema.FileResponse{
				Sha256:      sha256,
				Message:     err.Error(),
				Description: "파일을 찾을 수 없습니다",
			})
		}
	}
	err = json.Unmarshal(b, &file)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(http.StatusInternalServerError, schema.FileResponse{
				Sha256:      sha256,
				Message:     err.Error(),
				Description: "작업 중 에러가 발생하였습니다",
			})
		}
	}

	_, alreadyScanned := file.MultiAV["last_scan"]
	if alreadyScanned {
		_, ok := file.MultiAV["first_scan"]
		if !ok {
			file.MultiAV["first_scan"] = file.MultiAV["last_scan"]
		}
	}

	file, err = s.ms.FileUpdate(ctx, file)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(http.StatusInternalServerError, schema.FileResponse{
				Sha256:      sha256,
				Message:     err.Error(),
				Description: "업데이트 중 에러가 발생하였습니다",
			})
		}
	}
	return c.JSON(http.StatusOK, schema.FileResponse{
		Message:     "DB update succeed",
		Description: "DB 에 정상적으로 업데이트 완료",
	})
}

func (s *Server) UploadFile(c echo.Context) error {
	responseFile, err := c.FormFile("file")
	uploadDate := time.Now().UTC()
	if err != nil {
		return c.JSON(http.StatusBadRequest, schema.FileResponse{
			Message:     "파일이 없습니다",
			Description: "파일이 제대로 업로드 되었는지 확인해주세요",
		})
	}
	if responseFile.Size > s.Config.MaxFileSize {
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
	fileData, err := s.ms.FileSearch(c.Request().Context(), sha256)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(http.StatusInternalServerError, schema.FileResponse{
				Sha256:      sha256,
				Message:     err.Error(),
				FileName:    responseFile.Filename,
				FileSize:    responseFile.Size,
				Description: "작업 중 에러가 발생하였습니다",
			})
		}
	}
	if errors.Is(err, mongo.ErrNoDocuments) {
		ctx := context.Background()
		uploadedInfo, err := s.minio.Upload(ctx, responseFile)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, schema.FileResponse{
				Message:     "파일 업로드 실패",
				Description: err.Error(),
				FileName:    responseFile.Filename,
				FileSize:    responseFile.Size,
			})
		}
		uploadFile := schema.File{
			FileKey:          uploadedInfo.Key,
			Sha256:           sha256,
			FirstSubmission:  &uploadDate,
			LastSubmission:   &uploadDate,
			Size:             responseFile.Size,
			SubmissionsCount: 1,
		}
		submission := schema.Submission{
			Date:     &uploadDate,
			Filename: responseFile.Filename,
			Country:  "South Korea", // 추후 추가 예정
		}
		ch, err := s.rb.Channel()
		if err != nil {
			s.Logger.Error("rabbitmq channel error", err)
		}
		defer ch.Close()
		q, err := ch.QueueDeclare(s.rb.Config.FileScanQueue, false, false, false, false, nil)
		if err != nil {
			s.Logger.Error("rabbitmq queue error", err)
		}
		if err = ch.Publish("", q.Name, false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(fmt.Sprintf(`{"minio_object_key":"%s", "sha256":"%s"}`, uploadedInfo.Key, sha256)),
		}); err != nil {
			return c.JSON(http.StatusInternalServerError, schema.FileResponse{
				Message:     "분석 작업 요청 실패",
				Description: "Internal error",
				FileName:    responseFile.Filename,
				Sha256:      sha256,
			})
		}
		uploadFile.Submissions = append(uploadFile.Submissions, submission)
		_, err = s.ms.FileCreate(c.Request().Context(), uploadFile)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, schema.FileResponse{
				Sha256:      sha256,
				Message:     err.Error(),
				FileName:    responseFile.Filename,
				FileSize:    responseFile.Size,
				Description: "작업 중 에러가 발생하였습니다",
			})
		}
		return c.JSON(http.StatusOK, schema.FileResponse{
			Sha256:      sha256,
			Message:     "Success!",
			FileName:    responseFile.Filename,
			FileSize:    responseFile.Size,
			Description: "분석 작업이 요청이 성공적으로 제출되었습니다",
		})
	}
	sub := schema.Submission{
		Date:     &uploadDate,
		Filename: responseFile.Filename,
		Country:  "South Korea",
	}
	fileData.Submissions = append(fileData.Submissions, sub)
	fileData.LastSubmission = &uploadDate
	fileData.SubmissionsCount = fileData.SubmissionsCount + 1

	response, err := s.ms.FileUpdate(c.Request().Context(), fileData)
	if err != nil {
		return c.JSON(http.StatusOK, schema.FileResponse{
			Sha256:      sha256,
			Message:     err.Error(),
			FileName:    responseFile.Filename,
			FileSize:    responseFile.Size,
			Description: "작업 중 에러가 발생하였습니다",
		})
	}
	return c.JSON(http.StatusOK, response)
}
