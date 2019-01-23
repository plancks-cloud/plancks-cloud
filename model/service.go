package model

const ServiceCollectionName = "Service"

type Service struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
