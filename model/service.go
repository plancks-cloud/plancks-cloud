package model

const ServiceCollectionName = "services"

type Service struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
