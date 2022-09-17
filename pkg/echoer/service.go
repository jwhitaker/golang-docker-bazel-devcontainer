package echoer

import (
	log "github.com/sirupsen/logrus"
)

// Service is a struct to test how bazel deals with packages.
type Service struct {
}

// NewService will create a new Service instance
func NewService() Service {
	return Service{}
}

// GetMessage will return the message given
func (srv Service) GetMessage(message string) string {
	log.WithFields(log.Fields{
		"message": message,
	}).Debugln("Getting the message")

	return message
}
