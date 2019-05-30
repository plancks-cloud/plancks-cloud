package model

import "fmt"

const RouteCollectionName = "Route"

type Routes []Route

func (routes Routes) AnySSL() bool {
	for _, r := range routes {
		if r.SSL.Accept {
			return true
		}
	}
	return false
}

type Route struct {
	ID         string `json:"id"`
	DomainName string `json:"domainName"`
	Address    string `json:"address"`
	SSL        SSL    `json:"ssl"`
}

type SSL struct {
	Email  string `json:"email"`
	Accept bool   `json:"accept"`
}

func (r *Route) GetHttpAddress() string {
	return fmt.Sprint("http://", r.Address)
}

func (r *Route) GetWsAddress() string {
	return r.Address
}
