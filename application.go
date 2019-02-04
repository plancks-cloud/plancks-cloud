package main

import (
	"flag"
	"github.com/plancks-cloud/plancks-cloud/controller"
	"github.com/plancks-cloud/plancks-cloud/io/http-admin"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"github.com/plancks-cloud/plancks-cloud/model"
	"github.com/sirupsen/logrus"
)

var (
	addr       = flag.String("admin", ":6227", "TCP address to listen to")
	persistID  = flag.String("persistID", "", "Persistence ID")
	persistKey = flag.String("persistKey", "", "Persistence key")
	persistUrl = flag.String("persistURL", "...", "Persistence URL")
)

func main() {
	flag.Parse()
	logrus.Println("☁️☁️☁️ Planck's Cloud is starting ☁️☁️☁️")

	cred := getCreds()
	logrus.Println("...️ pulling down state")
	controller.StartupSync(cred)

	logrus.Println("...️ starting in-memory DB")
	mem.Init()

	logrus.Println("... ️starting health server")
	controller.StartHealthServer()

	logrus.Println("...️ starting api")
	http_admin.Startup(addr)

}

func getCreds() *model.Cred {
	return &model.Cred{
		URL: persistUrl,
		ID:  persistID,
		Key: persistKey,
	}
}
