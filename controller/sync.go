package controller

import (
	"encoding/json"
	"github.com/plancks-cloud/plancks-cloud/io/gcp"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/sirupsen/logrus"
)

func StartupSync(cred *model.Cred) {
	syncServicesDown(cred)
	syncRoutesDown(cred)
}

func syncRoutesDown(cred *model.Cred) {
	r, err := gcp.GetCollection(*cred.URL, *cred.ID, *cred.Key, model.RouteCollectionName)
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

func syncServicesDown(cred *model.Cred) {
	r, err := gcp.GetCollection(*cred.URL, *cred.ID, *cred.Key, model.ServiceCollectionName)
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
