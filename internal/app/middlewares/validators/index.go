package validators

import (
	"github.com/stetsd/monk-api/internal/app/constants"
	"github.com/stetsd/monk-api/internal/errorsApp"
	"github.com/stetsd/monk-api/internal/infrastructure/logger"
	"github.com/stetsd/monk-api/internal/tools/helpers"
	"net/http"
)

func handleErrorNoBody(err string, w http.ResponseWriter) {
	helpers.Throw(w, http.StatusInternalServerError, &constants.EmptyString)
	logger.Log.Error(
		errorsApp.Error(err),
	)
}

func handleErrorValidation(err error, w http.ResponseWriter, req *http.Request) {
	errorText := err.Error()
	logger.Log.ErrorHttp(req, errorText, http.StatusBadRequest)
	helpers.Throw(w, http.StatusBadRequest, &errorText)
}
