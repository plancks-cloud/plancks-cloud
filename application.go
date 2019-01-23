package main

import (
	"github.com/plancks-cloud/plancks-cloud/io/http-router"
	"log"
)

func main() {
	log.Println("☁️☁️☁️ Planck's Cloud is starting ☁️☁️☁️")

	//mem.Init()

	http_router.Proxy()

	//http_admin.Startup()

}
