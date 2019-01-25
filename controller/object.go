package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/plancks-cloud/plancks-cloud/model"
)

func HandleApply(item *model.Object) (err error) {
	if item.Type == "route" {
		return handleApplyRoutes(item.List)
	} else if item.Type == "service" {
		err = handleApplyServices(item.List)

		//TODO: Call docker client ensure services are there and up to day

	} else {
		err = errors.New(fmt.Sprint("Unknown type for /apply object: ", item.Type))
	}
	return
}

func handleApplyRoutes(list json.RawMessage) (err error) {
	var routes = &[]model.Route{}
	err = json.Unmarshal(list, routes)
	if err != nil {
		return
	}
	err = InsertManyRoutes(routes)
	if err != nil {
		return
	}
	RefreshProxy()
	return

}

func handleApplyServices(list json.RawMessage) (err error) {
	var s = &[]model.Service{}
	err = json.Unmarshal(list, s)
	if err != nil {
		return
	}
	err = InsertManyServices(s)
	if err != nil {
		return
	}
	return
}
