package domain

type Template struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Html   string `json:"html"`
}
