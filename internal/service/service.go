package service

import (
	"io"

	"github.com/chaikadn/file-converter/internal/models"
)

type CreateTaskRequest struct {
	File           io.Reader             `json:"-"`
	Filename       string                `json:"filename"`
	ConversionType models.ConversionType `json:"conversion_type"`
}

type ConversionService interface {
	CreateTask(req CreateTaskRequest) (*models.Task, error)

	GetTask(taskID string) (*models.Task, error)

	// ListTasks()
	// CancelTask()
}

type FileService interface {
	SaveUploadedFile(file io.Reader, filename string) (string, error)
	GetFile(filename string) (io.ReadCloser, error)
	DeleteFile(filename string) error
}
