package sources

import "enterpret/internal/shared"

type Service struct {
	repository    Repository
	twitterSource TwitterSource
}

func NewService(repository Repository, twitterSource TwitterSource) *Service {
	return &Service{repository: repository, twitterSource: twitterSource}
}

func (s *Service) PullData(startDate string, endDate string, key string, sourceId string, tenantSourceId string) ([]shared.Feedback, error) {
	source, err := s.repository.FindById(sourceId)
	if err != nil {
		return make([]shared.Feedback, 0), err
	}
	switch source.Type {
	case Twitter:
		rawData, err := s.twitterSource.PullData(startDate, endDate, key, source.Url)
		if err != nil {
			return make([]shared.Feedback, 0), err
		}
		return s.twitterSource.Transform(rawData, tenantSourceId)
	default:
		return make([]shared.Feedback, 0), err
	}
}
