package http_admin

import (
	"encoding/json"
	"flag"
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
	h := requestHandler
	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
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
	} else {
		log.Fatalln("Unhandled route! ", requestURI)
	}
	ctx.SetContentType("application/json; charset=utf8")

}

func handleService(method string, body []byte, ctx *fasthttp.RequestCtx) {
	var item *model.Service
	err := json.Unmarshal(body, item)
	if err != nil {
		//TODO: Return err
	}
	err = controller.Upsert(item)
	if err != nil {
		//TODO: Return err
	}
	ctx.Response.SetStatusCode(http.StatusOK)
	ctx.Response.SetBody(model.OKMessage)

}

func handleRoute(method string, body []byte, ctx *fasthttp.RequestCtx) {
	var item *model.Route
	err := json.Unmarshal(body, item)
	if err != nil {
		//TODO: Return err
	}
	err = controller.Upsert(item)
	if err != nil {
		//TODO: Return err
	}
	ctx.Response.SetStatusCode(http.StatusOK)
	ctx.Response.SetBody(model.OKMessage)

}
