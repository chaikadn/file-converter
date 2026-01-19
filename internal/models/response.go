package models

import "time"

type UploadResponse struct {
	ID     string     `json:"id"`
	Status TaskStatus `json:"status"`
}

type TaskStatusResponse struct {
	ID             string     `json:"id"`
	Status         TaskStatus `json:"status"`
	OriginalFormat string     `json:"original_format"`
	TargetFormat   string     `json:"target_format"`
	CreatedAt      time.Time  `json:"created_at"`
	ErrorMessage   string     `json:"error_message,omitempty"`
}
