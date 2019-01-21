package model

type InternalRoute struct {
	ID          string `json:"id"`
	ServiceName string `json:"serviceName"`
	DomainName  string `json:"domainName"`
}
