package controller

import (
	"encoding/json"
	"fmt"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func StartupSync(persistPath *string) {
	if len(*persistPath) == 0 {
		logrus.Info("Persist path not provided. Not starting persist")
		return
	}

	err := SaveConfig(&model.Config{ID: model.PersistPath, Val: *persistPath})
	if err != nil {
		logrus.Error(err)
		return
	}

	configPath := GetPersistPath()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = os.Mkdir(configPath, os.ModeDir)
		if err != nil {
			logrus.Error(err)
			return
		}
	}

	//syncServicesDown()
	syncRoutesDown()
}

func syncRoutesDown() {

	configPath := GetPersistPath()
	file := fmt.Sprint(configPath, "/", model.RouteCollectionName, ".json")
	if _, err := os.Stat(file); err == nil {
		b, err := ioutil.ReadFile(file)
		var arr []model.Route
		err = json.Unmarshal(b, &arr)
		if err != nil {
			logrus.Error(err)
			return
		}
		err = InsertManyRoutes(&arr)
		if err != nil {
			logrus.Error(err)
			return
		}
	}

}

//func syncRoutesUp() {
//	sl := GetAllRoutesCopy()
//	b, err := json.Marshal(sl)
//	if err != nil {
//		logrus.Error(err)
//		return
//	}
//	r, err := gcp.SetCollection(endpoint, id, key, model.RouteCollectionName, b)
//	if err != nil {
//		logrus.Error(err)
//		return
//	}
//	msg := string(r)
//	logrus.Println(msg)
//
//}
//
//func syncServicesDown() {
//	r, err := gcp.GetCollection(*cred.URL, *cred.ID, *cred.Key, model.ServiceCollectionName)
//	if err != nil {
//		logrus.Error(err)
//		return
//	}
//
//	var sl *[]model.Service
//	err = json.Unmarshal(r, sl)
//	if err != nil {
//		logrus.Error(err)
//		return
//	}
//
//	if len(*sl) > 0 {
//		err = InsertManyServices(sl)
//		if err != nil {
//			logrus.Error(err)
//			return
//		}
//	}
//
//}
//
//func syncServicesUp() {
//	sl := GetAllRoutesCopy()
//	b, err := json.Marshal(sl)
//	if err != nil {
//		logrus.Error(err)
//		return
//	}
//	r, err := gcp.SetCollection(endpoint, id, key, model.RouteCollectionName, b)
//	if err != nil {
//		logrus.Error(err)
//		return
//	}
//	msg := string(r)
//	logrus.Println(msg)
//
//}
