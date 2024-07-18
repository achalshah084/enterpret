package feedBacks

import (
	"enterpret/internal/shared"
	"enterpret/internal/sources"
	"enterpret/internal/tenantSources"
)

type Service struct {
	repository              Repository
	twitterSource           sources.TwitterSource
	tenantSourcesRepository tenantSources.Repository
	sourcesRepository       sources.Repository
}

func NewService(repository Repository, twitterSource sources.TwitterSource, tenantSourcesRepository tenantSources.Repository) *Service {
	return &Service{repository: repository, twitterSource: twitterSource, tenantSourcesRepository: tenantSourcesRepository}
}

func (s *Service) PushData(feedBackId string, tenantSourceId string) error {
	tenantSource, err := s.tenantSourcesRepository.FindById(tenantSourceId)
	if err != nil {
		return err
	}
	source, err := s.sourcesRepository.FindById(tenantSource.SourceId)
	if err != nil {
		return err
	}
	switch source.Type {
	case sources.Twitter:
		rawData, err := s.twitterSource.PullDataWithId(feedBackId, tenantSource.Key, source.Url)
		if err != nil {
			return err
		}
		data, err := s.twitterSource.TransformDataWithId(rawData, tenantSource.Id)
		if err != nil {
			return err
		}
		err = s.repository.Insert(data)
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}

func (s *Service) FindAll() []shared.Feedback {
	return s.repository.FindAll()
}
