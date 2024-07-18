package feedBacks

import (
	"encoding/json"
	"enterpret/internal/shared"
	"log"
	"net/http"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) PushFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	var pushRequest PushRequest
	if err := json.NewDecoder(r.Body).Decode(&pushRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if pushRequest.Id == "" || pushRequest.SourceId == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	err := h.Service.PushData(pushRequest.Id, pushRequest.SourceId)
	if err != nil {
		http.Error(w, "Error while pushing data. Please try again", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	feedBacks := h.Service.FindAll()
	json.NewEncoder(w).Encode(map[string][]shared.Feedback{"feedBacks": feedBacks})
}
