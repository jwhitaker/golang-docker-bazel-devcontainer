package app

import (
	log "github.com/sirupsen/logrus"
)

type Service struct {
}

func (s Service) GetMessage() string {
	log.Debugln("getting the message")

	return "Hello World"
}
