package main

import (
	"flag"
	"github.com/plancks-cloud/plancks-cloud/controller"
	"github.com/plancks-cloud/plancks-cloud/io/http-admin"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/sirupsen/logrus"
)

var (
	addr = flag.String("admin", ":6227", "TCP address to listen to")
)

func main() {
	flag.Parse()
	logrus.Println("☁️☁️☁️ Planck's Cloud is starting ☁️☁️☁️")

	logrus.Println("... ️starting health server")
	controller.StartHealthServer()

	logrus.Println("...️ starting in-memory DB")
	mem.Init()

	logrus.Println("...️ starting api")
	http_admin.Startup(addr)

}
