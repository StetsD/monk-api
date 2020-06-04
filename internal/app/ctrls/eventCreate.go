package ctrls

import (
	"encoding/json"
	"github.com/stetsd/monk-api/internal/app/constants"
	"github.com/stetsd/monk-api/internal/app/schemas"
	"github.com/stetsd/monk-api/internal/domain/services"
	"net/http"
)

func EventCreate(_ http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		serviceEvent, ok := req.Context().Value(services.ServiceEventName).(services.ServiceEvent)

		if !ok {
			handleErrorServiceUndefined("crtls.eventCreate: missed field \""+services.ServiceEventName+"\" in ctx", w)
			return
		}

		body := req.Context().Value(constants.BodyJson).(schemas.EventBody)

		eventResult, err := serviceEvent.GrpcConn.SendEvent(&body)

		if err != nil {
			handleErrorInternal(err.Error(), w)
			return
		}

		result := schemas.IdResult{Id: int(eventResult.EventId)}
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
