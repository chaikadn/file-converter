package server

import (
	"net/http"

	"github.com/chaikadn/file-converter/internal/service"
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
	w.Write([]byte("Upload files here\n"))
}

func (h *Handler) handleTaskStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Check task status here\n"))
}

func (h *Handler) handleDownload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Download files here\n"))
}
