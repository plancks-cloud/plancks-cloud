package http_admin

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/plancks-cloud/plancks-cloud/controller"
	"github.com/plancks-cloud/plancks-cloud/model"
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

	if requestURI == "/service" {
		handleService(method, ctx.Request.Body(), ctx)
	} else if requestURI == "/route" {
		handleRoute(method, ctx.Request.Body(), ctx)
	} else if requestURI == "/apply" {
		handleAny(method, ctx.Request.Body(), ctx)
	} else {
		log.Println("Unhandled route! ", requestURI)
	}
	ctx.SetContentType("application/json; charset=utf8")

}

func handleService(method string, body []byte, ctx *fasthttp.RequestCtx) {
	if method == http.MethodPost || method == http.MethodPut {
		var item = &model.Service{}
		fmt.Println(string(body))
		err := json.Unmarshal(body, item)
		if err != nil {
			fmt.Println(err)
			//TODO: respond with 500 to client
			return
		}
		err = controller.Upsert(item)
		if err != nil {
			fmt.Println(err)
			//TODO: respond with 500 to client
			return
		}
		ctx.Response.SetStatusCode(http.StatusOK)
		ctx.Response.SetBody(model.OKMessage)

	} else if method == http.MethodGet {
		ch := controller.GetAllServices()
		var arr []*model.Service
		for item := range ch {
			arr = append(arr, item)
		}
		b, err := json.Marshal(arr)
		if err != nil {
			fmt.Println(err)
			//TODO: respond with 500 to client
			return
		}
		//Send back empty array not null
		ctx.Response.SetStatusCode(http.StatusOK)
		if len(arr) == 0 {
			ctx.Response.SetBody([]byte("[]"))
			return
		}
		ctx.Response.SetBody(b)
	}

}

func handleRoute(method string, body []byte, ctx *fasthttp.RequestCtx) {
	if method == http.MethodPost || method == http.MethodPut {
		var item = &model.Route{}
		fmt.Println(string(body))
		err := json.Unmarshal(body, item)
		if err != nil {
			fmt.Println(err)
			//TODO: respond with 500 to client
			return
		}
		err = controller.Upsert(item)
		if err != nil {
			fmt.Println(err)
			//TODO: respond with 500 to client
			return
		}
		ctx.Response.SetStatusCode(http.StatusOK)
		ctx.Response.SetBody(model.OKMessage)
		controller.RefreshProxy()

	} else if method == http.MethodGet {
		ch := controller.GetAllRoutes()
		var arr []*model.Route
		for item := range ch {
			arr = append(arr, item)
		}
		b, err := json.Marshal(arr)
		if err != nil {
			fmt.Println(err)
			//TODO: respond with 500 to client
			return
		}
		//Send back empty array not null
		ctx.Response.SetStatusCode(http.StatusOK)
		if len(arr) == 0 {
			ctx.Response.SetBody([]byte("[]"))
			return
		}
		ctx.Response.SetBody(b)
	}

}

func handleAny(method string, body []byte, ctx *fasthttp.RequestCtx) {
	if method == http.MethodPost || method == http.MethodPut {
		var item = &model.Object{}
		err := json.Unmarshal(body, &item)
		if err != nil {
			fmt.Println(err)
			//TODO: http reply
			return
		}
		if item.Type == "route" {
			fmt.Println("Going to look for routes")
			var routes = &[]model.Route{}
			err := json.Unmarshal(item.List, routes)
			if err != nil {
				fmt.Println(err)
			}
			err = controller.InsertMany(routes)
			if err != nil {
				fmt.Println(err)
			}
			controller.RefreshProxy()

		} else if item.Type == "service" {
			panic("Service handler for /apply not implemented! ")
		} else {
			panic(fmt.Sprint("Unknown object type in apply: ", item.Type))
		}

		ctx.Response.SetStatusCode(http.StatusOK)
		ctx.Response.SetBody(model.OKMessage)

	}

}
