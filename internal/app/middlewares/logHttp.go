package middlewares

import (
	"github.com/stetsd/monk-api/internal/infrastructure/logger"
	"net/http"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		logger.Log.Info(req.Method + " " + req.RequestURI)

		next.ServeHTTP(w, req)
	})
}
