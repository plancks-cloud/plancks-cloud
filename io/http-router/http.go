package http_router

//Taken from from https://gist.github.com/bradfitz/1d7bdf12278d4d713212ce6c74875dab

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var (
	target   = "http://127.0.0.1:8000"
	targetWS = "127.0.0.1:8000"
)

func Proxy() {

	u, err := url.Parse(target)
	if err != nil {
		fmt.Println(err)
	}
	rp := httputil.NewSingleHostReverseProxy(u)

	httpsServer := &http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hj, isHJ := w.(http.Hijacker)
			if r.Header.Get("Upgrade") == "websocket" && isHJ {
				c, br, err := hj.Hijack()
				if err != nil {
					log.Printf("websocket websocket hijack: %v", err)
					http.Error(w, err.Error(), 500)
					return
				}
				defer c.Close()

				var be net.Conn
				be, err = net.DialTimeout("tcp", targetWS, 10*time.Second)

				if err != nil {
					log.Printf("websocket Dial: %v", err)
					http.Error(w, err.Error(), 500)
					return
				}
				defer be.Close()
				if err := r.Write(be); err != nil {
					log.Printf("websocket backend write request: %v", err)
					http.Error(w, err.Error(), 500)
					return
				}
				errc := make(chan error, 1)
				go func() {
					n, err := io.Copy(be, br) // backend <- buffered reader
					if err != nil {
						err = fmt.Errorf("websocket: to copy backend from buffered reader: %v, %v", n, err)
					}
					errc <- err
				}()
				go func() {
					n, err := io.Copy(c, be) // raw conn <- backend
					if err != nil {
						err = fmt.Errorf("websocket: to raw conn from backend: %v, %v", n, err)
					}
					errc <- err
				}()
				if err := <-errc; err != nil {
					log.Print(err)
				}
				return
			}
			rp.ServeHTTP(w, r)
		}),
	}

	panic(httpsServer.ListenAndServe())

}
