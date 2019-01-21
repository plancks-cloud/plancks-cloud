package model

var (
	Local *Cluster
)

type Cluster struct {
	Routes   *[]Route
	Services *[]Service
}
