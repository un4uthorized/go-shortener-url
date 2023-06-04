package entities

type URL struct {
	ID          string
	OriginalURL string
}

type ShortenURLRequest struct {
	URL string `json:"url" binding:"required"`
}
