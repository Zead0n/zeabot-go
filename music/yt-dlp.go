package music

type YtResult struct {
	Playlist  *string `json:"playlist"`
	Url       string  `json:"webpage_url"`
	Thumbnail string  `json:"thumbnail"`
}
