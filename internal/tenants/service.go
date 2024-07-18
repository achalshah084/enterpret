package tenants

import (
	"enterpret/internal/feedBacks"
	"enterpret/internal/tenantSources"
)

type Service struct {
	repository           Repository
	tenantSourcesService tenantSources.Service
	feedBacksRepository  feedBacks.Repository
}

func NewService(repository Repository, tenantSourcesService tenantSources.Service, feedBacksRepository feedBacks.Repository) *Service {
	return &Service{repository: repository, tenantSourcesService: tenantSourcesService, feedBacksRepository: feedBacksRepository}
}

func (s *Service) PullData(startDate string, endDate string, id string) error {
	tenant, err := s.repository.findById(id)
	if err != nil {
		return err
	}
	for _, tenantSourceId := range tenant.TenantSourceIds {
		response, err := s.tenantSourcesService.PullData(startDate, endDate, tenantSourceId)
		if err != nil {
			return err
		}
		for _, individualFeedback := range response {
			err := s.feedBacksRepository.Insert(individualFeedback)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
