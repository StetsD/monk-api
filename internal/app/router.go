package app

import (
	"github.com/gorilla/mux"
	"github.com/stetsd/monk-api/internal/app/constants"
	"github.com/stetsd/monk-api/internal/app/ctrls"
	"github.com/stetsd/monk-api/internal/app/middlewares"
	"github.com/stetsd/monk-api/internal/app/middlewares/validators"
	"github.com/stetsd/monk-api/internal/domain/services"
	"github.com/stetsd/monk-api/internal/tools"
)

func NewHttpRouter(serviceCollection *tools.ServiceCollection) *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.Log)

	if serviceCollection.ServiceUser != (services.ServiceUser{}) {
		// Create user
		routeRegistration := router.
			Path("/registration").
			Subrouter()
		routeRegistration.Methods("POST")
		routeRegistration.Use(
			middlewares.BodyParser(constants.RegistrationBody),
			validators.Registration,
			middlewares.ServiceCtxInjector(services.ServiceUserName, serviceCollection),
			ctrls.Registration,
		)
	}

	if serviceCollection.ServiceEvent != (services.ServiceEvent{}) {
		// Create event
		routeCreateEvent := router.
			Path("/event").
			Subrouter()
		routeCreateEvent.Methods("POST")
		routeCreateEvent.Use(
			middlewares.BodyParser(constants.EventBody),
			// TODO: check auth
			validators.EventCreate,
			middlewares.ServiceCtxInjector(services.ServiceEventName, serviceCollection),
			ctrls.EventCreate,
		)
	}

	return router
}
