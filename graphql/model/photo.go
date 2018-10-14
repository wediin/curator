package model

type Photo struct {
	ID          string   `json:"id"`
	Contributor string   `json:"contributor"`
	Urls        []string `json:"urls"`
	Timestamp   string   `json:"timestamp"`
	Masked      bool     `json:"masked"`
}
