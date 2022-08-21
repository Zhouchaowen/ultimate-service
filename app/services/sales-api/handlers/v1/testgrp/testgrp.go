package testgrp

import (
	"context"
	"github.com/Zhouchaowen/ultimate-service/foundation/web"
	"go.uber.org/zap"
	"net/http"
)

// Handlers manages the set of check endpoints.
type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handler is for development
func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	data := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, data, http.StatusOK)
}
