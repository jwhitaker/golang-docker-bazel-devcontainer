package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/thewhitakers/echoer/pkg/echoer"
)

func main() {
	svc := echoer.NewService()

	log.Println(svc.GetMessage("Hello World"))
}
