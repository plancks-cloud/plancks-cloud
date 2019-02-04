package controller

import (
	"encoding/json"
	"github.com/plancks-cloud/plancks-cloud/io/gcp"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/sirupsen/logrus"
)

func StartupSync(api, id, key string) {
	syncServicesDown(api, id, key)
	syncRoutesDown(api, id, key)
}

func syncRoutesDown(endpoint, id, key string) {
	r, err := gcp.GetCollection(endpoint, id, key, model.RouteCollectionName)
	if err != nil {
		logrus.Error(err)
		return
	}

	var sl *[]model.Route
	err = json.Unmarshal(r, sl)
	if err != nil {
		logrus.Error(err)
		return
	}

	if len(*sl) > 0 {
		err = InsertManyRoutes(sl)
		if err != nil {
			logrus.Error(err)
			return
		}
	}

}

func syncRoutesUp(endpoint, id, key string) {
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

func syncServicesDown(endpoint, id, key string) {
	r, err := gcp.GetCollection(endpoint, id, key, model.ServiceCollectionName)
	if err != nil {
		logrus.Error(err)
		return
	}

	var sl *[]model.Service
	err = json.Unmarshal(r, sl)
	if err != nil {
		logrus.Error(err)
		return
	}

	if len(*sl) > 0 {
		err = InsertManyServices(sl)
		if err != nil {
			logrus.Error(err)
			return
		}
	}

}

func syncServicesUp(endpoint, id, key string) {
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
