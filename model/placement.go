package model

type Placement struct {
	ID        string `json:"id"`
	ServiceID string `json:"serviceId"`
	ClusterID string `json:"clusterId"`
}
