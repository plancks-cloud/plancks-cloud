package controller

import (
	"flag"
	"github.com/hashicorp/go-memdb"
	"github.com/plancks-cloud/plancks-cloud/io/http-router"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/sirupsen/logrus"
)

var (
	proxy = flag.String("proxy", ":6228", "TCP address to listen to")
	stop  chan bool
)

func GetAllRoutes() (resp chan model.Route) {
	resp = make(chan model.Route)
	go func() {
		ite, err := mem.GetAll(model.RouteCollectionName)
		iteratorToManyRoutes(ite, err, resp)
		close(resp)
	}()
	return resp
}

func GetAllRoutesCopy() (result model.Routes) {
	result = []model.Route{}
	for item := range GetAllRoutes() {
		result = append(result, item)
	}
	return
}

func InsertManyRoutes(routes *[]model.Route) (err error) {
	for _, route := range *routes {
		cRoute := route //Seems redundant - it's not. Pointers be crazy
		err = mem.Push(&cRoute)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	syncRoutesToDisk()
	return
}

func iteratorToManyRoutes(iterator memdb.ResultIterator, err error, out chan model.Route) {
	iteratorToHandler(iterator, err, func(next interface{}) {
		item := next.(*model.Route)
		out <- *item
	})
}

func RefreshProxy() {
	arr := GetAllRoutesCopy()
	http_router.StopServer(stop)
	stop = http_router.Serve(*proxy, arr)
}

func DeleteManyRoutes(routes *[]model.Route) (err error) {
	for _, route := range *routes {
		_, err = mem.Delete(model.RouteCollectionName, model.RouteCollectionID, route.ID)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	syncRoutesToDisk()
	return
}
