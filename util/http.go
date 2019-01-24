package util

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/http"
	"strings"
)

func HostOfURL(url string) string {
	//TODO: proper handling with errors
	sl := strings.Split(url, ":")
	return sl[0]
}

func WriteJsonResponseToReq(ctx *fasthttp.RequestCtx, code int, resp []byte) {
	ctx.Response.SetStatusCode(code)
	ctx.Response.SetBody(resp)
}

func WriteErrorToReq(ctx *fasthttp.RequestCtx, msg string) {
	ctx.Response.SetStatusCode(http.StatusInternalServerError)
	var r struct {
		OK  bool   `json:"ok"`
		Msg string `json:"msg"`
	}
	r.OK = false
	r.Msg = msg
	b, _ := json.Marshal(&r)
	ctx.Response.SetBody(b)
}
