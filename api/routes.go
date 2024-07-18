package api

import (
	"enterpret/internal/feedBacks"
	"enterpret/internal/tenants"
	"github.com/gorilla/mux"
)

// RegisterRoutes registers the routes for the HTTP server
func RegisterRoutes(tenantHandler *tenants.Handler, feedBacksHandler *feedBacks.Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/pull", tenantHandler.PullFeedbackHandler).Methods("GET")
	router.HandleFunc("/push", feedBacksHandler.PushFeedbackHandler).Methods("POST")

	router.HandleFunc("/findAll", feedBacksHandler.FindAll).Methods("GET")

	return router
}
