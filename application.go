package main

import (
	"flag"
	"github.com/plancks-cloud/plancks-cloud/io/http-admin"
	"github.com/plancks-cloud/plancks-cloud/io/mem"
	"log"
)

func main() {
	flag.Parse()
	log.Println("☁️☁️☁️ Planck's Cloud is starting ☁️☁️☁️")

	mem.Init()
	http_admin.Startup()

}
