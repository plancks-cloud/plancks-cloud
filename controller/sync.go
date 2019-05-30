package controller

import (
	"encoding/json"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

var PersistPath string

func StartupSync(persistPath string) {
	if persistPath == "" {
		logrus.Info("Persist path not provided. Not starting persist")
		return
	}

	syncServicesFromDisk()
	syncRoutesFromDisk()
}

//Saves routes to memory DB
func syncRoutesFromDisk() {

	if PersistPath == "" {
		return
	}

	file := filepath.ToSlash(filepath.Join(PersistPath, model.RouteCollectionFileName))
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
		RefreshProxy()
	} else {
		logrus.Println("Routes json not found. Not loading any routes.")
	}

}

//Saves routes to memory DB
func syncServicesFromDisk() {

	if PersistPath == "" {
		return
	}

	file := filepath.ToSlash(filepath.Join(PersistPath, model.ServiceCollectionFileName))
	if _, err := os.Stat(file); err == nil {
		b, err := ioutil.ReadFile(file)
		var arr []model.Service
		err = json.Unmarshal(b, &arr)
		if err != nil {
			logrus.Error(err)
			return
		}
		err = InsertManyServices(&arr)
		if err != nil {
			logrus.Error(err)
			return
		}
	} else {
		logrus.Println("Services json not found. Not loading any routes.")
	}

}

//Saves routes to disk
func syncRoutesToDisk() {

	if PersistPath == "" {
		return
	}
	file := filepath.ToSlash(filepath.Join(PersistPath, model.RouteCollectionFileName))
	logrus.Println("Saving routes to:", file)
	//Get routes -> json -> []byte
	routes := GetAllRoutesCopy()
	b, err := json.Marshal(&routes)
	if err != nil {
		logrus.Error(err)
		return
	}

	//Delete file
	os.Remove(file)

	//Save file
	err = ioutil.WriteFile(file, b, 0644)
	if err != nil {
		logrus.Error(err)
	}

}

//Saves svcs to disk
func syncServicesToDisk() {

	if PersistPath == "" {
		return
	}

	file := filepath.ToSlash(filepath.Join(PersistPath, model.ServiceCollectionFileName))
	logrus.Println("Saving services to:", file)
	//Get services -> json -> []byte
	arr := GetAllServicesCopy()
	b, err := json.Marshal(&arr)
	if err != nil {
		logrus.Error(err)
		return
	}

	//Delete file
	os.Remove(file)

	//Save file
	err = ioutil.WriteFile(file, b, 0644)
	if err != nil {
		logrus.Error(err)
	}

}
