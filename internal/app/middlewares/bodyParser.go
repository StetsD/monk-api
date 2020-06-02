package middlewares

import (
	"context"
	"encoding/json"
	"github.com/stetsd/monk-api/internal/app/constants"
	"github.com/stetsd/monk-api/internal/app/schemas"
	"github.com/stetsd/monk-api/internal/infrastructure/logger"
	"github.com/stetsd/monk-api/internal/tools/helpers"
	"io/ioutil"
	"net/http"
)

func BodyParser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		bodyRaw, err := ioutil.ReadAll(req.Body)

		if err != nil {
			logger.Log.Error(err.Error())
			helpers.Throw(w, http.StatusInternalServerError, &constants.EmptyString)
			return
		}

		var msg schemas.RegistrationBody
		err = json.Unmarshal(bodyRaw, &msg)

		if err != nil {
			logger.Log.ErrorHttp(req, err.Error(), http.StatusBadRequest)
			helpers.Throw(w, http.StatusBadRequest, &constants.EmptyString)
			return
		}

		ctx := context.WithValue(req.Context(), constants.BodyJson, msg)
		req = req.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
