package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/plancks-cloud/plancks-docker/controller/pc-docker"
	"github.com/sirupsen/logrus"
)

func HandleApply(item *model.Object) (err error) {
	if item.Type == "route" {
		return handleApplyRoutes(item.List)
	} else if item.Type == "service" {
		err = handleApplyServices(item.List)
	} else {
		err = errors.New(fmt.Sprint("Unknown type for /apply object: ", item.Type))
	}
	return
}

func RawToRoutes(list json.RawMessage) (routes *[]model.Route, err error) {
	routes = &[]model.Route{}
	err = json.Unmarshal(list, routes)
	return
}

func handleApplyRoutes(list json.RawMessage) (err error) {
	routes, err := RawToRoutes(list)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infoln("Inserting ", len(*routes), " routes")
	err = InsertManyRoutes(routes)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infoln("Refreshing proxy")
	RefreshProxy()
	return

}

func rawToServices(list json.RawMessage) (routes *[]model.Service, err error) {
	routes = &[]model.Service{}
	err = json.Unmarshal(list, routes)
	return
}

func handleApplyServices(list json.RawMessage) (err error) {
	services, err := rawToServices(list)
	if err != nil {
		logrus.Error(err)
		return
	}
	err = InsertManyServices(services)
	if err != nil {
		logrus.Error(err)
		return
	}
	healthDoorbell <- true //Ensures the health check runs
	return
}

func HandleDelete(item *model.Object) (err error) {
	if item.Type == "route" {
		return handleDeleteRoutes(item.List)
	} else if item.Type == "service" {
		err = handleDeleteServices(item.List)

		//TODO: Call docker client ensure services are there and up to day

	} else {
		err = errors.New(fmt.Sprint("Unknown type for /apply object: ", item.Type))
	}
	return
}

func handleDeleteRoutes(list json.RawMessage) (err error) {
	routes, err := RawToRoutes(list)
	if err != nil {
		logrus.Error(err)
		return
	}
	err = DeleteManyRoutes(routes)
	if err != nil {
		logrus.Error(err)
		return
	}
	RefreshProxy()
	return

}

func handleDeleteServices(list json.RawMessage) (err error) {
	services, err := rawToServices(list)
	if err != nil {
		logrus.Error(err)
		return
	}
	err = DeleteManyServices(services)
	if err != nil {
		logrus.Error(err)
		return
	}
	err = pc_docker.DeleteServices(convertServices(services))
	if err != nil {
		logrus.Error("Failed to delete service: ", err)
	}
	healthDoorbell <- true //Ensures the health check runs
	return
}
