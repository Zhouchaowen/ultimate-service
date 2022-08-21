package mid

import (
	"context"
	"github.com/Zhouchaowen/ultimate-service/foundation/web"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// Logger ...
func Logger(log *zap.SugaredLogger) web.Middleware {
	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			traceId := "000000000000000000"
			statusCode := http.StatusOK
			now := time.Now()

			log.Infow("request started", "traceid", traceId, "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr)

			// Call the next handler.
			err := handler(ctx, w, r)

			log.Infow("request completed", "traceid", traceId, "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr, "statuscode", statusCode, "since", time.Since(now))

			// Return the error so it can be handled further up the chain.
			return err
		}

		return h
	}
	return m
}
