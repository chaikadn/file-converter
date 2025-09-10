package server

import "net/http"

type handler struct {
}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) handleUpload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Upload files here\n"))
}

func (h *handler) handleTaskStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Check task status here\n"))
}

func (h *handler) handleDownload(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Download files here\n"))
}
