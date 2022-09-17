package main

import (
	log "github.com/sirupsen/logrus"
	// "log"

	"github.com/thewhitakers/echoer/pkg/app"
)

func main() {
	srv := app.Service{}

	log.Println(srv.GetMessage())
}
