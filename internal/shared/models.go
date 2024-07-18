package shared

type Feedback interface {
	GetID() string
	GetType() string
	GetContent() string
	GetMetadata() Metadata
}

type Metadata interface {
	GetTenantSourceId() string
	GetLanguage() string
}

type BaseFeedback struct {
	Id       string       `json:"id"`
	Content  string       `json:"content"`
	Type     FeedBackType `json:"type"`
	Metadata Metadata     `json:"metadata"`
}

type Review struct {
	BaseFeedback
	Rating int `json:"rating"`
}

type Conversation struct {
	BaseFeedback
	Participants []string `json:"participants"`
}

func (review Review) GetID() string {
	return review.Id
}

func (review Review) GetType() string {
	return review.Id
}

func (review Review) GetContent() string {
	return review.Id
}

func (review Review) GetMetadata() Metadata {
	return review.Metadata
}

func (conversation Conversation) GetID() string {
	return conversation.Id
}

func (conversation Conversation) GetType() string {
	return conversation.Id
}

func (conversation Conversation) GetContent() string {
	return conversation.Id
}

func (conversation Conversation) GetMetadata() Metadata {
	return conversation.Metadata
}

func (conversation Conversation) GetParticipants() []string {
	return conversation.Participants
}

type TwitterMetadata struct {
	Language       string `json:"language"`
	TenantSourceId string `json:"tenantSourceId"`
	Country        string `json:"country"`
}

func (twitterMetadata TwitterMetadata) GetTenantSourceId() string {
	return twitterMetadata.TenantSourceId
}

func (twitterMetadata TwitterMetadata) GetLanguage() string {
	return twitterMetadata.Language
}

func (twitterMetadata TwitterMetadata) GetCountry() string {
	return twitterMetadata.Country
}
