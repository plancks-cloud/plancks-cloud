package model

type Route struct {
	ID        string `json:"id"`
	ServiceID string `json:"serviceId"`
	DomainID  string `json:"domainID"`
}
