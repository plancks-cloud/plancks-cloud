package model

const ServiceCollectionName = "Service"
const ServiceCollectionID = "id"

type Service struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Replicas    int    `json:"replicas,omitempty"`
	MemoryLimit int    `json:"memoryLimit,omitempty"`
	Network     string `json:"network"`
}
