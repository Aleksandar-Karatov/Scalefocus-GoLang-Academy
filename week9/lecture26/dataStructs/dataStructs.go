package dataStructs

type NewsID = int

type IDPayload struct {
	IDs []NewsID
}

type Story struct {
	Title string `json:"title"`
	Score int    `json:"score"`
}
type Stories struct {
	AllStories []Story `json:"top_stories"`
}

type SimpleHandler struct{}

type StoriesPageData struct {
	PageTitle string
	Links     []Story
}
