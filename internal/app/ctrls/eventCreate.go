package ctrls

import (
	"encoding/json"
	"fmt"
	"github.com/stetsd/monk-api/internal/app/constants"
	"github.com/stetsd/monk-api/internal/app/schemas"
	"github.com/stetsd/monk-api/internal/infrastructure/logger"
	"github.com/stetsd/monk-api/internal/tools/helpers"
	"net/http"
)

func EventCreate(_ http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		body := req.Context().Value(constants.BodyJson).(schemas.EventBody)

		fmt.Println(body)

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
