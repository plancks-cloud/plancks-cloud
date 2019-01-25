package main

import (
	"flag"
	"github.com/plancks-cloud/plancks-cloud/controller"
	"github.com/plancks-cloud/plancks-cloud/io/http-admin"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"log"
)

var (
	addr = flag.String("admin", ":6227", "TCP address to listen to")
)

func main() {
	flag.Parse()
	log.Println("☁️☁️☁️ Planck's Cloud is starting ☁️☁️☁️")

	controller.StartHealthServer()

	mem.Init()
	http_admin.Startup(addr)

}
