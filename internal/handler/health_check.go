package handler

import (
	"context"
	"net/http"
)

var bodyHealthCheck = []byte("ok")

func HealthCheck(_ context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(bodyHealthCheck)
}
