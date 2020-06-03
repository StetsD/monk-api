package ctrls

import (
	"github.com/stetsd/monk-api/internal/app/constants"
	"github.com/stetsd/monk-api/internal/errorsApp"
	"github.com/stetsd/monk-api/internal/infrastructure/logger"
	"github.com/stetsd/monk-api/internal/tools/helpers"
	"net/http"
)

func handleErrorServiceUndefined(err string, w http.ResponseWriter) {
	logger.Log.Error(
		errorsApp.Error(err),
	)
	helpers.Throw(w, http.StatusInternalServerError, &constants.EmptyString)
}

func handleErrorInternal(err string, w http.ResponseWriter) {
	logger.Log.Error(err)
	helpers.Throw(w, http.StatusInternalServerError, &constants.EmptyString)
}

func handleErrorBadRequest(errorText *string, w http.ResponseWriter, req *http.Request) {
	logger.Log.ErrorHttp(req, *errorText, http.StatusBadRequest)
	helpers.Throw(w, http.StatusBadRequest, errorText)
}
