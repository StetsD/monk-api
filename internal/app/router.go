package app

import (
	"github.com/gorilla/mux"
	"github.com/stetsd/monk-api/internal/app/ctrls"
	"github.com/stetsd/monk-api/internal/app/middlewares"
	"github.com/stetsd/monk-api/internal/app/middlewares/validators"
	"github.com/stetsd/monk-api/internal/domain/services"
	"github.com/stetsd/monk-api/internal/tools"
)

func NewHttpRouter(serviceCollection *tools.ServiceCollection) *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.Log)

	if &serviceCollection.ServiceUser != nil {
		// Registration user
		routeRegistration := router.
			Path("/registration").
			Subrouter()
		routeRegistration.Methods("POST")
		routeRegistration.Use(
			middlewares.BodyParser,
			validators.Registration,
			middlewares.ServiceCtxInjector(services.ServiceUserName, serviceCollection),
			ctrls.Registration,
		)
	}

	return router
}
