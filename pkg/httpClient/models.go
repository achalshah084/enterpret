package httpClient

type TwitterFeedback struct {
	Id           string   `json:"id"`
	Content      string   `json:"content"`
	Type         string   `json:"type"`
	Country      string   `json:"country,omitempty"`
	Language     string   `json:"language,omitempty"`
	UserId       string   `json:"userId,omitempty"`
	Participants []string `json:"participants,omitempty"` // Assuming twitter returns this
	Rating       int      `json:"rating,omitempty"`       // Assuming twitter also has review system and return rating for a post
}
