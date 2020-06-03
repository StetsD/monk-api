package ctrls

import (
	"encoding/json"
	"github.com/stetsd/monk-api/internal/app/constants"
	"github.com/stetsd/monk-api/internal/app/schemas"
	"github.com/stetsd/monk-api/internal/domain/errorsDomain"
	"github.com/stetsd/monk-api/internal/domain/services"
	"github.com/stetsd/monk-api/internal/errorsApp"
	"github.com/stetsd/monk-api/internal/infrastructure/logger"
	"github.com/stetsd/monk-api/internal/tools/helpers"
	"net/http"
)

func Registration(_ http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		serviceUser, ok := req.Context().Value(services.ServiceUserName).(services.ServiceUser)

		if !ok {
			logger.Log.Error(
				errorsApp.Error("crtls.registration: missed field \"" + services.ServiceUserName + "\" in ctx"),
			)
			helpers.Throw(w, http.StatusInternalServerError, &constants.EmptyString)
			return
		}

		body := req.Context().Value(constants.BodyJson).(schemas.RegistrationBody)

		err := serviceUser.Registration(&body)

		if err != nil {
			_, ok := err.(errorsDomain.ErrorUser)
			if ok {
				errorText := err.Error()
				logger.Log.ErrorHttp(req, errorText, http.StatusBadRequest)
				helpers.Throw(w, http.StatusBadRequest, &errorText)
			} else {
				logger.Log.Error(err.Error())
				helpers.Throw(w, http.StatusInternalServerError, &constants.EmptyString)
			}
			return
		}

		result := schemas.HttpResult{Result: "success"}
		jsonBody, err := json.Marshal(result)

		if err != nil {
			logger.Log.Error(err.Error())
			helpers.Throw(w, http.StatusInternalServerError, &constants.EmptyString)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonBody)

		if err != nil {
			logger.Log.Error(err.Error())
			helpers.Throw(w, http.StatusInternalServerError, &constants.EmptyString)
			return
		}

	})
}
