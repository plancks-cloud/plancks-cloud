package main

import (
	"github.com/plancks-cloud/plancks-cloud/io/http-admin"
	"log"
)

func main() {
	log.Println("☁️☁️☁️ Planck's Cloud is starting ☁️☁️☁️")

	http_admin.Startup()

}
