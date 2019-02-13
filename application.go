package main

import (
	"flag"
	"github.com/plancks-cloud/plancks-cloud/controller"
	"github.com/plancks-cloud/plancks-cloud/io/http-admin"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/sirupsen/logrus"
)

var (
	addr        = flag.String("admin", ":6227", "TCP address to listen to")
	persistPath = flag.String("persistPath", "C:\\tmp", "Persistence path")
)

func main() {
	flag.Parse()
	logrus.Println("☁️☁️☁️ Planck's Cloud is starting ☁️☁️☁️")

	logrus.Println("...️ starting in-memory DB")
	mem.Init()

	logrus.Println("...️ pulling down state")
	controller.StartupSync(persistPath)

	logrus.Println("... ️starting health server")
	controller.StartHealthServer()

	logrus.Println("...️ starting api")
	http_admin.Startup(addr)

}
