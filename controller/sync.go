package controller

import (
	"encoding/json"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
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

	syncServicesFromDisk()
	syncRoutesFromDisk()
}

//Saves routes to memory DB
func syncRoutesFromDisk() {

	file := filepath.ToSlash(filepath.Join(GetPersistPath(), model.RouteCollectionFileName))
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
	} else {
		logrus.Println("Routes json not found. Not loading any routes.")
	}

}

//Saves routes to memory DB
func syncServicesFromDisk() {

	file := filepath.ToSlash(filepath.Join(GetPersistPath(), model.ServiceCollectionFileName))
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

	//Check if feature is on
	c := GetConfig(model.PersistPath)
	if c.ID == "" || c.Val == "" {
		logrus.Infoln("No persist path. No syncing.")
		return
	}

	//Get routes -> json -> []byte
	routes := GetAllRoutesCopy()
	b, err := json.Marshal(&routes)
	if err != nil {
		logrus.Error(err)
		return
	}

	file := filepath.ToSlash(filepath.Join(GetPersistPath(), model.RouteCollectionFileName))

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

	//Check if feature is on
	c := GetConfig(model.PersistPath)
	if c.ID == "" || c.Val == "" {
		logrus.Infoln("No persist path. No syncing.")
		return
	}

	//Get services -> json -> []byte
	arr := GetAllServicesCopy()
	b, err := json.Marshal(&arr)
	if err != nil {
		logrus.Error(err)
		return
	}

	file := filepath.ToSlash(filepath.Join(GetPersistPath(), model.ServiceCollectionFileName))

	//Delete file
	os.Remove(file)

	//Save file
	err = ioutil.WriteFile(file, b, 0644)
	if err != nil {
		logrus.Error(err)
	}

}
