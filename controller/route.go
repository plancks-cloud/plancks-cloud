package controller

import (
	"flag"
	"fmt"
	"github.com/hashicorp/go-memdb"
	"github.com/plancks-cloud/plancks-cloud/io/http-router"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/sirupsen/logrus"
)

var (
	proxy = flag.String("proxy", ":80", "TCP address to listen to")
	stop  chan bool
)

func GetAllRoutes() (resp chan *model.Route) {
	resp = make(chan *model.Route)
	go func() {
		ite, err := mem.GetAll(model.RouteCollectionName)
		iteratorToManyRoutes(ite, err, resp)
		close(resp)
	}()
	return resp
}

func GetAllRoutesCopy() []model.Route {
	var arr []model.Route
	for item := range GetAllRoutes() {
		arr = append(arr, *item)
	}
	return arr
}

func InsertManyRoutes(routes *[]model.Route) (err error) {
	for _, route := range *routes {
		logrus.Infoln("Upserting", route.ID, route.DomainName)
		err = Upsert(&route)
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	return
}

func iteratorToManyRoutes(iterator memdb.ResultIterator, err error, out chan *model.Route) {
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	if iterator == nil {
		return
	}
	more := true
	count := 0
	for more {
		next := iterator.Next()
		if next == nil {
			more = false
			continue
		}
		item := next.(*model.Route)
		out <- item
		count++
	}
	logrus.Debugln(fmt.Sprintf("Route iterator counts: %d", count))
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
	return
}
