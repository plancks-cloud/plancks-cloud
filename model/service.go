package model

type Service struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"name"`
	Replicas int    `json:"replicas"`
}
