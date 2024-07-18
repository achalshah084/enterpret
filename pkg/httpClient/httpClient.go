package httpClient

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) PullTwitterData(startDate string, endDate string, key string, url string) ([]TwitterFeedback, error) {
	return []TwitterFeedback{
		{
			Id:      "123",
			Content: "very well said",
			Type:    "conversation",
			Country: "india",
			UserId:  "789",
		},
		{
			Id:      "123",
			Content: "very well said",
			Type:    "conversation",
			Country: "india",
			UserId:  "789",
		},
		{
			Id:      "234",
			Content: "very well said",
			Type:    "conversation",
			Country: "usa",
			UserId:  "890",
		},
	}, nil
}

func (s *Service) PullTwitterDataWithId(id string, key string, url string) (TwitterFeedback, error) {
	return TwitterFeedback{
		Id:      "123",
		Content: "very well said again",
		Type:    "conversation",
		Country: "india",
		UserId:  "789",
	}, nil
}
