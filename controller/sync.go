package controller

import (
	"encoding/json"
	"github.com/plancks-cloud/plancks-cloud/io/gcp"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/sirupsen/logrus"
)

func syncRoutes(endpoint, id, key string) {
	sl := GetAllRoutesCopy()
	b, err := json.Marshal(sl)
	if err != nil {
		logrus.Error(err)
		return
	}
	r, err := gcp.SetCollection(endpoint, id, key, model.RouteCollectionName, b)
	if err != nil {
		logrus.Error(err)
		return
	}
	msg := string(r)
	logrus.Println(msg)

}

func syncServices(endpoint, id, key string) {
	sl := GetAllRoutesCopy()
	b, err := json.Marshal(sl)
	if err != nil {
		logrus.Error(err)
		return
	}
	r, err := gcp.SetCollection(endpoint, id, key, model.RouteCollectionName, b)
	if err != nil {
		logrus.Error(err)
		return
	}
	msg := string(r)
	logrus.Println(msg)

}
