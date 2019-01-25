package http_router

//Taken from from https://gist.github.com/bradfitz/1d7bdf12278d4d713212ce6c74875dab

import (
	"bufio"
	"fmt"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/plancks-cloud/plancks-cloud/util"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func Serve(listenAddr string, prev chan bool, routes []model.Route) chan bool {
	if prev != nil {
		fmt.Println("Stopping proxy server")
		prev <- true
		time.Sleep(250 * time.Millisecond) //Not sure how necessary this is...
	}
	return startProxy(listenAddr, routes)
}

func startProxy(listenAddr string, routes []model.Route) (stop chan bool) {
	fmt.Println("Starting proxy server")
	stop = make(chan bool)

	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Println(err)
	}

	m := newReverseProxyMap(routes)
	go func() {
		_ = http.Serve(l, newReverseProxyHandler(routes, m))
	}()
	go func() {
		<-stop
		err = l.Close()
		if err != nil {
			log.Fatal(err)
		}

	}()
	return

}

func newReverseProxyMap(routes []model.Route) map[string]*httputil.ReverseProxy {
	m := make(map[string]*httputil.ReverseProxy)
	for _, route := range routes {
		u, err := url.Parse(route.GetHttpAddress())
		if err != nil {
			fmt.Println(err)
			//TODO: check this before it gets here?
		}
		m[route.DomainName] = httputil.NewSingleHostReverseProxy(u)
	}
	return m

}

func newReverseProxyHandler(routes []model.Route, m map[string]*httputil.ReverseProxy) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hj, isHJ := w.(http.Hijacker)
		if r.Header.Get("Upgrade") == "websocket" && isHJ {
			handleHiJackedWS(hj, r, w, routes)
			return
		}
		rp := m[util.HostOfURL(r.Host)]
		if rp == nil {
			fmt.Println("Could not find host: ", util.HostOfURL(r.Host))
			//TODO: Send error
			return
		}
		rp.ServeHTTP(w, r)
	})
}

func handleHiJackedWS(hj http.Hijacker, r *http.Request, w http.ResponseWriter, routes []model.Route) {
	c, br, err := hj.Hijack()
	if err != nil {
		log.Printf("websocket websocket hijack: %v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	defer c.Close()

	var be net.Conn
	found, route := findRouteByHost(routes, util.HostOfURL(r.Host))
	if !found {
		fmt.Println("Could not find domain name among routes: ", util.HostOfURL(r.Host))
		return
	}
	be, err = net.DialTimeout("tcp", route.GetWsAddress(), 10*time.Second)

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
	startWSReadWrite(&be, br, errc, &c)
	if err := <-errc; err != nil {
		log.Print(err)
	}
}

func startWSReadWrite(be *net.Conn, br *bufio.ReadWriter, errc chan error, c *net.Conn) {
	go func() {
		n, err := io.Copy(*be, br) // backend <- buffered reader
		if err != nil {
			err = fmt.Errorf("websocket: to copy backend from buffered reader: %v, %v", n, err)
		}
		errc <- err
	}()
	go func() {
		n, err := io.Copy(*c, *be) // raw conn <- backend
		if err != nil {
			err = fmt.Errorf("websocket: to raw conn from backend: %v, %v", n, err)
		}
		errc <- err
	}()
}

func findRouteByHost(routes []model.Route, host string) (found bool, route model.Route) {
	for _, route = range routes {
		if route.DomainName == host {
			found = true
			return
		}
	}
	return
}
