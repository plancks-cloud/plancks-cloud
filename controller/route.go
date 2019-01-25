package controller

import (
	"flag"
	"github.com/hashicorp/go-memdb"
	"github.com/plancks-cloud/plancks-cloud/io/http-router"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/plancks-cloud/plancks-cloud/model"
)

var (
	proxy = flag.String("proxy", ":6228", "TCP address to listen to")
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
		err = Upsert(&route)
		if err != nil {
			return err
		}
	}
	return
}

func iteratorToManyRoutes(iterator memdb.ResultIterator, err error, out chan *model.Route) {
	c := mem.IteratorToChannel(iterator, err)
	for i := range c {
		item := i.(*model.Route)
		out <- item
	}

}

func RefreshProxy() {
	arr := GetAllRoutesCopy()
	http_router.StopServer(stop)
	stop = http_router.Serve(*proxy, arr)
}
