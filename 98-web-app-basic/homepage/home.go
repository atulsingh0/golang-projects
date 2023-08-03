package homepage

import (
	"log"
	"net/http"
)

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("request processed")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("Welcome! to GO World"))
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
