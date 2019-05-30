package http_admin

import (
	"encoding/json"
	"fmt"
	"github.com/plancks-cloud/plancks-cloud/controller"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/plancks-cloud/plancks-cloud/util"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/valyala/fasthttp"
)

var handlers = make(map[string]ApplyHandler)

type ApplyHandler func(method string, body []byte, ctx *fasthttp.RequestCtx)

func Startup(addr *string) {
	setupHandlers()
	if err := fasthttp.ListenAndServe(*addr, requestHandler); err != nil {
		logrus.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func setupHandlers() {
	handlers["/apply"] = handleApply
	handlers["/delete"] = handleDelete
	handlers["/service"] = handleService
	handlers["/route"] = handleRoute
}

func requestHandler(ctx *fasthttp.RequestCtx) {

	method := string(ctx.Method())
	requestURI := string(ctx.Request.RequestURI())

	f, ok := handlers[requestURI]
	if ok {
		f(method, ctx.Request.Body(), ctx)
	} else {
		logrus.Println("Unhandled route! ", requestURI)
		util.WriteErrorToReq(ctx, fmt.Sprint("Could not find a route for ", requestURI))
	}

}

func handleService(method string, body []byte, ctx *fasthttp.RequestCtx) {
	var arr []*model.Service
	for item := range controller.GetAllServices() {
		arr = append(arr, item)
	}
	b, err := json.Marshal(arr)
	if err != nil {
		logrus.Println(err)
		util.WriteErrorToReq(ctx, fmt.Sprint(err.Error()))
		return
	}
	//Send back empty array not null
	if len(arr) == 0 {
		b = []byte("[]")
	}
	util.WriteJsonResponseToReq(ctx, http.StatusOK, b)
}

func handleRoute(method string, body []byte, ctx *fasthttp.RequestCtx) {
	arr := controller.GetAllRoutesCopy()
	b, err := json.Marshal(arr)
	if err != nil {
		logrus.Println(err)
		util.WriteErrorToReq(ctx, fmt.Sprint(err.Error()))
		return
	}
	if len(arr) == 0 {
		b = []byte("[]")
	}
	util.WriteJsonResponseToReq(ctx, http.StatusOK, b)

}

func handleApply(method string, body []byte, ctx *fasthttp.RequestCtx) {
	if method == http.MethodPost || method == http.MethodPut {
		var item = &model.Object{}
		err := json.Unmarshal(body, &item)
		if err != nil {
			logrus.Println(err)
			util.WriteErrorToReq(ctx, fmt.Sprint(err.Error()))
			return
		}

		err = controller.HandleApply(item)
		if err != nil {
			logrus.Println(err)
			util.WriteErrorToReq(ctx, fmt.Sprint(err.Error()))
			return
		}

		ctx.Response.SetStatusCode(http.StatusOK)
		ctx.Response.Header.Add("Content-type", "application/json")
		ctx.Response.SetBody(model.OKMessage)
	}
}

func handleDelete(method string, body []byte, ctx *fasthttp.RequestCtx) {
	if method == http.MethodPost || method == http.MethodPut {
		var item = &model.Object{}
		err := json.Unmarshal(body, &item)
		if err != nil {
			logrus.Println(err)
			util.WriteErrorToReq(ctx, fmt.Sprint(err.Error()))
			return
		}

		err = controller.HandleDelete(item)
		if err != nil {
			logrus.Println(err)
			util.WriteErrorToReq(ctx, fmt.Sprint(err.Error()))
			return
		}

		ctx.Response.SetStatusCode(http.StatusOK)
		ctx.Response.Header.Add("Content-type", "application/json")
		ctx.Response.SetBody(model.OKMessage)

	}

}
