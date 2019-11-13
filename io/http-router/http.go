package http_router

//Taken from from https://gist.github.com/bradfitz/1d7bdf12278d4d713212ce6c74875dab

import (
	"bufio"
	"fmt"
	"github.com/mholt/certmagic"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/plancks-cloud/plancks-cloud/util"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func StopServer(prev chan bool) {
	if prev != nil {
		logrus.Println("Stopping proxy server...")
		prev <- true
		time.Sleep(50 * time.Millisecond) //Not sure how necessary this is...
	}
}

func Serve(listenAddr string, routes model.Routes) (stop chan bool) {

	//SERIOUS PROBLEMS HERE

	logrus.Println("Starting proxy server")
	stop = make(chan bool)

	if len(routes) == 0 {
		logrus.Println("No routes - not going to listen.")
		return
	}
	var magic *certmagic.Config
	email, hosts := describeSSL(routes)
	if len(hosts) > 0 {
		magic = makeMagic(email)
	}

	//HTTP traffic
	listenHTTP, err := net.Listen("tcp", listenAddr)
	if err != nil {
		logrus.Println(err)
	}

	m := newReverseProxyMap(routes)
	go func() {
		_ = http.Serve(listenHTTP, newReverseProxyHandler(routes, m, magic, false))
	}()

	if !routes.AnySSL() {
		logrus.Println("No SSL routes. Not listening.")
		return
	}

	certmagic.HTTPSPort = 6229
	listenTLS, err := certmagic.Listen(hosts)

	if err = magic.Manage(hosts); err != nil {
		logrus.Println(fmt.Errorf(err.Error()))
	}

	go func() {
		if err = http.Serve(listenTLS, newReverseProxyHandler(routes, m, magic, true)); err != nil {
			logrus.Println(fmt.Errorf(err.Error()))
		}
	}()
	startStopper(stop, listenHTTP, listenTLS)
	return

}

func startStopper(stop chan bool, listenHTTP net.Listener, listenTLS net.Listener) {
	var err error
	go func() {
		<-stop
		if err = listenHTTP.Close(); err != nil {
			logrus.Error(err)
		}
		if err = listenTLS.Close(); err != nil {
			logrus.Error(err)
		}
		close(stop)
	}()
}

func makeMagic(email string) *certmagic.Config {
	config := certmagic.Config{CA: certmagic.LetsEncryptProductionCA, Email: email, Agreed: true}
	cache := certmagic.NewCache(certmagic.CacheOptions{
		GetConfigForCert: func(certificate certmagic.Certificate) (c certmagic.Config, e error) {
			c = config
			return
		},
		OCSPCheckInterval:  7 * 24 * time.Hour,
		RenewCheckInterval: 7 * 24 * time.Hour,
	})
	return certmagic.New(cache, config)

}

func describeSSL(routes []model.Route) (email string, hosts []string) {
	logrus.Debugln("describeSSL received ", len(routes), " routes")
	for _, r := range routes {
		if r.SSL.Accept {
			email = r.SSL.Email
			hosts = append(hosts, r.DomainName)
		}
	}
	logrus.Debugln("describeSSL returning ", len(hosts), " routes")
	return
}

func newReverseProxyMap(routes []model.Route) map[string]*httputil.ReverseProxy {
	logrus.Debugln("newReverseProxyMap received ", len(routes), " routes")
	m := make(map[string]*httputil.ReverseProxy)
	for _, route := range routes {
		u, err := url.Parse(route.GetHttpAddress())
		if err != nil {
			logrus.Println(err)
			//TODO: check this before it gets here?
		}
		logrus.Debugln("newReverseProxyMap setting", route.DomainName)
		m[route.DomainName] = httputil.NewSingleHostReverseProxy(u)
	}
	logrus.Debugln("newReverseProxyMap returning", len(m), "rps")
	return m

}

func newReverseProxyHandler(routes model.Routes, m map[string]*httputil.ReverseProxy, magic *certmagic.Config, fromTLS bool) http.Handler {
	logrus.Debugln("Hosts known before: ", len(m))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if doNotCarryOn := magic.HandleHTTPChallenge(w, r); doNotCarryOn == true {
			return // challenge handled; nothing else to do
		}
		if carryOn := handleRedirect(w, r, routes, fromTLS); carryOn == false {
			return //handled by http redirect; nothing else to do
		}
		if carryOn := handleWS(w, r, routes); carryOn == false {
			return //handled by handleWS; nothing else to do
		}
		if carryOn := handleByReverseProxy(w, r, m, fromTLS); carryOn == false {
			return //Handled by reverse proxy; nothing else to do
		}
		logrus.Errorln("Could not handle: ", r.URL.String(), "among", len(m))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))

	})
}

func handleByReverseProxy(w http.ResponseWriter, r *http.Request, m map[string]*httputil.ReverseProxy, fromTLS bool) (carryOn bool) {
	rp, ok := m[util.HostOfURL(r.Host)]
	if ok {
		if fromTLS {
			r.Header.Add("X-Forwarded-Proto", "https")
		}
		rp.ServeHTTP(w, r)
		return false
	}
	return true
}

func handleRedirect(w http.ResponseWriter, r *http.Request, routes model.Routes, fromTLS bool) (carryOn bool) {
	found, route := routes.Find(r.Host)
	if found && !fromTLS && route.SSL.Accept && !route.AllowHTTP {
		result := "https://" + r.Host + r.URL.Path
		if len(r.URL.Query()) > 0 {
			result += "?" + r.URL.Query().Encode()
		}
		w.Header().Set("Location", result)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return false
	}
	return true

}

func handleWS(w http.ResponseWriter, r *http.Request, routes model.Routes) (carryOn bool) {
	hj, isHJ := w.(http.Hijacker)
	if r.Header.Get("Upgrade") == "websocket" && isHJ {
		handleHiJackedWS(hj, r, w, routes)
		return false
	}
	return true
}

func handleHiJackedWS(hj http.Hijacker, r *http.Request, w http.ResponseWriter, routes []model.Route) {
	c, br, err := hj.Hijack()
	if err != nil {
		logrus.Printf("websocket websocket hijack: %v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	defer c.Close()

	var be net.Conn
	found, route := findRouteByHost(routes, util.HostOfURL(r.Host))
	if !found {
		logrus.Println("Could not find domain name among routes: ", util.HostOfURL(r.Host))
		return
	}
	be, err = net.DialTimeout("tcp", route.GetWsAddress(), 10*time.Second)

	if err != nil {
		logrus.Printf("websocket Dial: %v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	defer be.Close()
	if err := r.Write(be); err != nil {
		logrus.Printf("websocket backend write request: %v", err)
		http.Error(w, err.Error(), 500)
		return
	}
	errc := make(chan error, 1)
	startWSReadWrite(&be, br, errc, &c)
	if err := <-errc; err != nil {
		logrus.Print(err)
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
