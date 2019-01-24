package http_admin

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/plancks-cloud/plancks-cloud/controller"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/plancks-cloud/plancks-cloud/util"
	"log"
	"net/http"

	"github.com/valyala/fasthttp"
)

var (
	addr = flag.String("admin", ":6227", "TCP address to listen to")
)

func Startup() {
	flag.Parse()
	if err := fasthttp.ListenAndServe(*addr, requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	log.Println(string(ctx.Method()))
	log.Println(string(ctx.Request.RequestURI()))

	method := string(ctx.Method())
	requestURI := string(ctx.Request.RequestURI())

	if requestURI == "/apply" {
		handleAny(method, ctx.Request.Body(), ctx)
	} else if requestURI == "/service" {
		handleService(method, ctx.Request.Body(), ctx)
	} else if requestURI == "/route" {
		handleRoute(method, ctx.Request.Body(), ctx)
	} else {
		log.Println("Unhandled route! ", requestURI)
	}
	util.WriteErrorToReq(ctx, fmt.Sprint("Could not find a route for ", requestURI))

}

func handleService(method string, body []byte, ctx *fasthttp.RequestCtx) {
	if method == http.MethodGet {
		ch := controller.GetAllServices()
		var arr []*model.Service
		for item := range ch {
			arr = append(arr, item)
		}
		b, err := json.Marshal(arr)
		if err != nil {
			fmt.Println(err)
			util.WriteErrorToReq(ctx, fmt.Sprint(err.Error()))
			return
		}
		//Send back empty array not null
		if len(arr) == 0 {
			b = []byte("[]")
		}
		util.WriteJsonResponseToReq(ctx, http.StatusOK, b)
	}

}

func handleRoute(method string, body []byte, ctx *fasthttp.RequestCtx) {
	if method == http.MethodGet {
		ch := controller.GetAllRoutes()
		var arr []*model.Route
		for item := range ch {
			arr = append(arr, item)
		}
		b, err := json.Marshal(arr)
		if err != nil {
			fmt.Println(err)
			util.WriteErrorToReq(ctx, fmt.Sprint(err.Error()))
			return
		}
		//Send back empty array not null
		if len(arr) == 0 {
			b = []byte("[]")
		}
		util.WriteJsonResponseToReq(ctx, http.StatusOK, b)
	}
}

func handleAny(method string, body []byte, ctx *fasthttp.RequestCtx) {
	if method == http.MethodPost || method == http.MethodPut {
		var item = &model.Object{}
		err := json.Unmarshal(body, &item)
		if err != nil {
			fmt.Println(err)
			util.WriteErrorToReq(ctx, fmt.Sprint(err.Error()))
			return
		}
		if item.Type == "route" {
			var routes = &[]model.Route{}
			err := json.Unmarshal(item.List, routes)
			if err != nil {
				fmt.Println(err)
				util.WriteErrorToReq(ctx, fmt.Sprint(err.Error()))
				return
			}
			err = controller.InsertManyRoutes(routes)
			if err != nil {
				fmt.Println(err)
			}
			controller.RefreshProxy()

		} else if item.Type == "service" {
			var s = &[]model.Service{}
			err := json.Unmarshal(item.List, s)
			if err != nil {
				fmt.Println(err)
				util.WriteErrorToReq(ctx, fmt.Sprint(err.Error()))
				return
			}
			err = controller.InsertManyServices(s)
			if err != nil {
				fmt.Println(err)
				util.WriteErrorToReq(ctx, fmt.Sprint(err.Error()))
				return
			}

			//TODO: Call docker client ensure services are there and up to day

		} else {
			util.WriteErrorToReq(ctx, fmt.Sprint("Could not handle type in object: ", item.Type))
			return
		}

		ctx.Response.SetStatusCode(http.StatusOK)
		ctx.Response.SetBody(model.OKMessage)

	}

}
