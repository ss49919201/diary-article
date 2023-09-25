package handler

import (
	"net/http"
)

var (
	bodyHealthCheck = []byte("ok")
)

type HealthCheck interface {
	Do(http.ResponseWriter, *http.Request)
}

type healthCheck struct{}

func NewHealthCheck() (HealthCheck, error) {
	return healthCheck{}, nil
}

func (_ healthCheck) Do(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(bodyHealthCheck)
}
