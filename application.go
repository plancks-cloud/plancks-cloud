package main

import (
	"github.com/plancks-cloud/plancks-cloud/io/http-router"
	"log"
	"time"
)

func main() {
	log.Println("☁️☁️☁️ Planck's Cloud is starting ☁️☁️☁️")

	//mem.Init()

	stopRP := http_router.StartProxy()

	//http_admin.Startup()

	time.Sleep(60 * time.Second)
	stopRP <- true

}
