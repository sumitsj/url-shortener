package contracts

type ShortenUrlResponse struct {
	OriginalUrl  string `json:",omitempty"`
	ShortenedUrl string `json:",omitempty"`
	Error        string `json:",omitempty"`
}
