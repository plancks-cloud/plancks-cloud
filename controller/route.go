package controller

import "github.com/plancks-cloud/plancks-cloud/model"

func GetAllRoutes() *[]*model.Route {
	return model.Routes
}
