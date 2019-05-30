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
		if err != nil {
			logrus.Errorln(err)
			close(resp)
			return
		}
		iteratorToManyRoutes(ite, resp)
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
		if err = mem.Push(&cRoute); err != nil {
			logrus.Error(err)
			return
		}
	}
	syncRoutesToDisk()
	return
}

func iteratorToManyRoutes(iterator memdb.ResultIterator, out chan model.Route) {
	iteratorToHandler(iterator, func(next interface{}) {
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
		if _, err = mem.Delete(model.RouteCollectionName, model.RouteCollectionID, route.ID); err != nil {
			logrus.Error(err)
			return err
		}
	}
	syncRoutesToDisk()
	return
}
