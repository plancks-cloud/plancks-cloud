package http_router

//Taken from from https://gist.github.com/bradfitz/1d7bdf12278d4d713212ce6c74875dab

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

var (
	hostPorts = []string{":10800"}
	base      = Base{":8080"}
)

type Base struct {
	Host string
}

func Proxy() {

	tlsConfig := &tls.Config{}

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
				if len(hostPorts) == 0 {
					backendHostPort := base.Host
					if !strings.Contains(backendHostPort, ":") {
						backendHostPort = net.JoinHostPort(backendHostPort, "443")
					}
					be, err = tls.DialWithDialer(dialer, "tcp", backendHostPort, tlsConfig)
				} else {
					for _, hostPort := range hostPorts {
						be, err = tls.DialWithDialer(dialer, "tcp", hostPort, tlsConfig)
						if err == nil {
							break
						}
					}
				}
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

	httpsServer.ListenAndServe()

}
