package main

import (
	"enterpret/api"
	"enterpret/internal/feedBacks"
	"enterpret/internal/sources"
	"enterpret/internal/tenantSources"
	"enterpret/internal/tenants"
	"enterpret/pkg/httpClient"
	"log"
	"net/http"
)

func main() {
	feedBacksRepository := feedBacks.NewRepository()
	sourcesRepository := sources.NewRepository()
	tenantsRepository := tenants.NewRepository()
	tenantSourcesRepository := tenantSources.NewRepository()

	httpClientService := *httpClient.NewService()
	twitterSource := sources.NewTwitterSource(httpClientService)

	feedBacksService := feedBacks.NewService(*feedBacksRepository, *twitterSource, *tenantSourcesRepository)
	sourcesService := sources.NewService(*sourcesRepository, *twitterSource)
	tenantSourcesService := tenantSources.NewService(*tenantSourcesRepository, *sourcesService)
	tenantsService := tenants.NewService(*tenantsRepository, *tenantSourcesService, *feedBacksRepository)

	tenantHandler := tenants.NewHandler(tenantsService)
	feedBackHandler := feedBacks.NewHandler(feedBacksService)

	router := api.RegisterRoutes(tenantHandler, feedBackHandler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
