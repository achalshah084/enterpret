package feedBacks

import "enterpret/internal/shared"

var globalFeedBackMap = make(map[string]shared.Feedback)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (repository Repository) Insert(feedBack shared.Feedback) error {
	// You need to perform find one and update in DB
	globalFeedBackMap[feedBack.GetID()] = feedBack
	return nil
}

func (repository Repository) FindAll() []shared.Feedback {
	feedBacks := make([]shared.Feedback, 0)
	for _, feedBack := range globalFeedBackMap {
		feedBacks = append(feedBacks, feedBack)
	}
	return feedBacks
}
