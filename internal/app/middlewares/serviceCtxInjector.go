package middlewares

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/stetsd/monk-api/internal/domain/services"
	"github.com/stetsd/monk-api/internal/tools"
	"net/http"
)

func ServiceCtxInjector(serviceName string, coll *tools.ServiceCollection) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var ctx context.Context

			switch serviceName {
			case services.ServiceUserName:
				ctx = context.WithValue(req.Context(), serviceName, coll.ServiceUser)
			}

			req = req.WithContext(ctx)

			next.ServeHTTP(w, req)
		})
	}

}
