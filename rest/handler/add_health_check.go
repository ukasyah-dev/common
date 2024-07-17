package handler

import (
	"context"
	"net/http"

	"github.com/ukasyah-dev/common/model"
	"github.com/ukasyah-dev/common/rest/server"
)

type Health struct {
	Status string `json:"status" example:"ok"`
}

func HealthCheck(ctx context.Context, req *model.Empty) (*Health, error) {
	return &Health{Status: "ok"}, nil
}

func AddHealthCheck(srv *server.Server) {
	Add(srv, http.MethodGet, "/", HealthCheck, Config{
		Summary:     "Health check",
		Description: "Check whether the server is ready to serve.",
		Tags:        []string{"Health"},
	})
}
