package model

import "fmt"

const RouteCollectionName = "Route"

type Route struct {
	ID         string `json:"id"`
	DomainName string `json:"domainName"`
	Address    string `json:"address"`
}

func (r *Route) GetHttpAddress() string {
	return fmt.Sprint("http://", r.Address)
}

func (r *Route) GetWsAddress() string {
	return r.Address
}
