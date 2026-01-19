package models

import "time"

type Task struct {
	ID             string         `json:"id"`
	Status         TaskStatus     `json:"status"`
	ConversionType ConversionType `json:"conversion_type"`
	CreatedAt      time.Time      `json:"created_at"`
	ErrorMessage   string         `json:"error_message,omitempty"`
	OriginalFile   string         `json:"original_file"`
	ConvertedFile  string         `json:"converted_file,omitempty"`
}

type ConversionType string

const (
	ConversionPNGtoJPEG ConversionType = "png_to_jpeg"
	ConversionJPEGtoPNG ConversionType = "jpeg_to_png"
)

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusProcessing TaskStatus = "processing"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusFailed     TaskStatus = "failed"
)
