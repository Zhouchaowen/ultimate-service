package testgrp

import (
	"context"
	"github.com/Zhouchaowen/ultimate-service/foundation/web"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
)

// Handlers manages the set of check endpoints.
type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handler is for development
func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// test error middleware
	if n := rand.Intn(100); n%2 == 0 {
		//return errors.New("untrusted error")
		//return validate.NewRequestError(errors.New("trusted error"),http.StatusBadRequest)
		return web.NewShutdownError("restart service")
	}

	data := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, data, http.StatusOK)
}
