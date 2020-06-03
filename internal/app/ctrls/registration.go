package ctrls

import (
	"encoding/json"
	"github.com/stetsd/monk-api/internal/app/constants"
	"github.com/stetsd/monk-api/internal/app/schemas"
	"github.com/stetsd/monk-api/internal/domain/errorsDomain"
	"github.com/stetsd/monk-api/internal/domain/services"
	"net/http"
)

func Registration(_ http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		serviceUser, ok := req.Context().Value(services.ServiceUserName).(services.ServiceUser)

		if !ok {
			handleErrorServiceUndefined("crtls.registration: missed field \""+services.ServiceUserName+"\" in ctx", w)
			return
		}

		body := req.Context().Value(constants.BodyJson).(schemas.RegistrationBody)

		err := serviceUser.Registration(&body)

		if err != nil {
			_, ok := err.(errorsDomain.ErrorUser)
			if ok {
				errorText := err.Error()
				handleErrorBadRequest(&errorText, w, req)
			} else {
				handleErrorInternal(err.Error(), w)
			}
			return
		}

		result := schemas.HttpResult{Result: "success"}
		jsonBody, err := json.Marshal(result)

		if err != nil {
			handleErrorInternal(err.Error(), w)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonBody)

		if err != nil {
			handleErrorInternal(err.Error(), w)
			return
		}

	})
}
