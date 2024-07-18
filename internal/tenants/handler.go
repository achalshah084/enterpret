package tenants

import (
	"encoding/json"
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

func (h *Handler) PullFeedbackHandler(w http.ResponseWriter, r *http.Request) {
	tenantID := r.URL.Query().Get("id")
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")
	if tenantID == "" || startDate == "" || endDate == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	err := h.Service.PullData(startDate, endDate, tenantID)
	if err != nil {
		http.Error(w, "Error while pulling data. Please try again", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}
