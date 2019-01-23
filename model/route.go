package model

const RouteCollectionName = "routes"

type Route struct {
	ID         string `json:"id"`
	DomainName string `json:"domainName"`
	Address    string `json:"address"`
}
