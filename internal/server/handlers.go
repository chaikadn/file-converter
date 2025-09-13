package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/chaikadn/file-converter/internal/models"
	"github.com/chaikadn/file-converter/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	fileService       service.FileService
	conversionService service.ConversionService
}

func NewHandler(fileService service.FileService, conversionService service.ConversionService) *Handler {
	return &Handler{
		fileService:       fileService,
		conversionService: conversionService,
	}
}

func (h *Handler) handleUpload(w http.ResponseWriter, r *http.Request) {
	contentType := strings.Split(r.Header.Get("Content-Type"), ";")[0]
	if contentType != "multipart/form-data" {
		h.respondError(w, http.StatusBadRequest, "wrong content type", fmt.Errorf("wrong content type: %s", contentType))
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10МБ
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "failed to parse form", fmt.Errorf("failed to parse form: %v", err))
		return
	}

	filename := r.FormValue("file_name")
	if filename == "" {
		h.respondError(w, http.StatusBadRequest, "empty filename", fmt.Errorf("empty filename"))
		return
	}

	targetFormat := r.FormValue("target_format")
	if targetFormat == "" {
		// сделать формат по умолчанию?
		h.respondError(w, http.StatusBadRequest, "empty target format", fmt.Errorf("empty target format"))
		return
	}

	originalFormat := strings.TrimPrefix(filepath.Ext(filename), ".")

	var conversionType models.ConversionType
	switch {
	case originalFormat == "png" && targetFormat == "jpeg":
		conversionType = models.ConversionPNGtoJPEG
	case originalFormat == "jpeg" && targetFormat == "png":
		conversionType = models.ConversionJPEGtoPNG
	default:
		h.respondError(
			w, http.StatusBadRequest,
			"unknown conversion format",
			fmt.Errorf("unknown conversion format: %s to %s", originalFormat, targetFormat),
		)
		return
	}

	mockTask := models.Task{
		ID:             uuid.NewString(),
		Status:         models.TaskStatusPending,
		ConversionType: conversionType,
		CreatedAt:      time.Now(),
		OriginalFile:   filename,
	}

	resp := models.UploadResponse{
		ID:     mockTask.ID,
		Status: mockTask.Status,
	}

	slog.Info(
		"task created",
		"id", mockTask.ID,
		"filename", mockTask.OriginalFile,
		"conversion_type", mockTask.ConversionType,
	)
	h.respondJSON(w, http.StatusAccepted, resp)

	// Тут скачать и ассинхронно обработать файл - добавить в очередь на конвертацию
}

func (h *Handler) handleTaskStatus(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")

	// temp
	mockTask := models.Task{
		ID:             "mock-task",
		Status:         models.TaskStatusPending,
		ConversionType: models.ConversionPNGtoJPEG,
		CreatedAt:      time.Now(),
		ErrorMessage:   "",
	}
	formats := strings.Split(string(mockTask.ConversionType), "_to_")

	if taskID != mockTask.ID {
		h.respondError(w, http.StatusNotFound, "task not found", fmt.Errorf("task id '%s' not found", taskID))
		return
	}

	resp := models.TaskStatusResponse{
		ID:             mockTask.ID,
		Status:         mockTask.Status,
		OriginalFormat: formats[0],
		TargetFormat:   formats[1],
		CreatedAt:      mockTask.CreatedAt,
		ErrorMessage:   mockTask.ErrorMessage,
	}

	h.respondJSON(w, http.StatusOK, resp)
}

func (h *Handler) handleDownload(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")

	w.Write([]byte("Download file " + taskID + "\n"))
}

func (h *Handler) respondError(w http.ResponseWriter, status int, msg string, err error) {
	slog.Error(msg, "error", err)
	h.respondJSON(w, status, map[string]string{"error": msg})
}

func (h *Handler) respondJSON(w http.ResponseWriter, status int, data any) {
	body, err := json.Marshal(data)
	if err != nil {
		slog.Error("failed to encode response body", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	body = append(body, '\n')

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	w.Write(body)
}
