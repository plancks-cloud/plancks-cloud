package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/plancks-cloud/plancks-cloud/model"
)

func HandleApply(item *model.Object) (err error) {
	if item.Type == "route" {
		var routes = &[]model.Route{}
		err = json.Unmarshal(item.List, routes)
		if err != nil {
			return
		}
		err = InsertManyRoutes(routes)
		if err != nil {
			return
		}
		RefreshProxy()

	} else if item.Type == "service" {
		var s = &[]model.Service{}
		err = json.Unmarshal(item.List, s)
		if err != nil {
			return
		}
		err = InsertManyServices(s)
		if err != nil {
			return
		}

		//TODO: Call docker client ensure services are there and up to day

	} else {
		err = errors.New(fmt.Sprint("Unknown type for /apply object: ", item.Type))
	}
	return
}
