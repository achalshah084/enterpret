package tenantSources

import (
	"enterpret/internal/shared"
	"enterpret/internal/sources"
)

type Service struct {
	repository     Repository
	sourcesService sources.Service
}

func NewService(repository Repository, sourcesService sources.Service) *Service {
	return &Service{repository: repository, sourcesService: sourcesService}
}

func (s *Service) PullData(startDate string, endDate string, id string) ([]shared.Feedback, error) {
	tenantSource, err := s.repository.FindById(id)
	if err != nil {
		return make([]shared.Feedback, 0), err
	}
	return s.sourcesService.PullData(startDate, endDate, tenantSource.Key, tenantSource.SourceId, tenantSource.Id)
}
