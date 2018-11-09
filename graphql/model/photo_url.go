package model

type PhotoURL struct {
	Width  int32  `json:"width"`
	Height int32  `json:"height"`
	URL    string `json:"url"`
}
