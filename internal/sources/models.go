package sources

import (
	"enterpret/internal/shared"
	"enterpret/pkg/httpClient"
)

type Sources struct {
	Id   string     `json:"id"`
	Type SourceType `json:"type"`
	Url  string     `json:"url"`
}

type ISources interface {
	PullData(startDate string, endDate string, key string, url string) (interface{}, error)
	Transform(response interface{}, tenantSourceId string) ([]shared.Feedback, error)
	PullDataWithId(id string, key string, url string) (interface{}, error)
	TransformDataWithId(response interface{}, tenantSourceId string) (shared.Feedback, error)
}

type TwitterSource struct {
	httpClient httpClient.Service
}

func (twitterSource TwitterSource) PullData(startDate string, endDate string, key string, url string) (interface{}, error) {
	twitterData, err := twitterSource.httpClient.PullTwitterData(startDate, endDate, key, url)
	if err != nil {
		return twitterData, err
	}
	return twitterData, nil
}

func (twitterSource TwitterSource) PullDataWithId(id string, key string, url string) (interface{}, error) {
	twitterData, err := twitterSource.httpClient.PullTwitterDataWithId(id, key, url)
	if err != nil {
		return twitterData, err
	}
	return twitterData, nil
}

func (twitterSource TwitterSource) TransformDataWithId(response interface{}, tenantSourceId string) (shared.Feedback, error) {
	feedback := response.(httpClient.TwitterFeedback)
	switch feedback.Type {
	case string(shared.ConversationFeedBackType):
		return shared.Conversation{
			BaseFeedback: shared.BaseFeedback{
				Id:      feedback.Id,
				Content: feedback.Content,
				Type:    shared.ConversationFeedBackType,
				Metadata: shared.TwitterMetadata{
					Language:       feedback.Language,
					TenantSourceId: tenantSourceId,
					Country:        feedback.Country,
				},
			},
			Participants: feedback.Participants,
		}, nil
	default:
		return shared.Review{
			BaseFeedback: shared.BaseFeedback{
				Id:      feedback.Id,
				Content: feedback.Content,
				Type:    shared.ConversationFeedBackType,
				Metadata: shared.TwitterMetadata{
					Language:       feedback.Language,
					TenantSourceId: tenantSourceId,
					Country:        feedback.Country,
				},
			},
			Rating: feedback.Rating,
		}, nil
	}
}

func NewTwitterSource(httpClient httpClient.Service) *TwitterSource {
	return &TwitterSource{httpClient: httpClient}
}

func (twitterSource TwitterSource) Transform(response interface{}, tenantSourceId string) ([]shared.Feedback, error) {
	twitterResponse := response.([]httpClient.TwitterFeedback)
	feedBackResponse := make([]shared.Feedback, 0)
	for _, feedback := range twitterResponse {
		switch feedback.Type {
		case string(shared.ConversationFeedBackType):
			feedBackResponse = append(feedBackResponse, shared.Conversation{
				BaseFeedback: shared.BaseFeedback{
					Id:      feedback.Id,
					Content: feedback.Content,
					Type:    shared.ConversationFeedBackType,
					Metadata: shared.TwitterMetadata{
						Language:       feedback.Language,
						TenantSourceId: tenantSourceId,
						Country:        feedback.Country,
					},
				},
				Participants: feedback.Participants,
			})
		default:
			feedBackResponse = append(feedBackResponse, shared.Review{
				BaseFeedback: shared.BaseFeedback{
					Id:      feedback.Id,
					Content: feedback.Content,
					Type:    shared.ConversationFeedBackType,
					Metadata: shared.TwitterMetadata{
						Language:       feedback.Language,
						TenantSourceId: tenantSourceId,
						Country:        feedback.Country,
					},
				},
				Rating: feedback.Rating,
			})
		}
	}
	return feedBackResponse, nil
}
