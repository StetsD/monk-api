package validators

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/stetsd/monk-api/internal/app/constants"
	"github.com/stetsd/monk-api/internal/app/schemas"
	"github.com/stetsd/monk-api/internal/errorsApp"
	"github.com/stetsd/monk-api/internal/infrastructure/logger"
	"github.com/stetsd/monk-api/internal/tools/helpers"
	"net/http"
)

func Registration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body := req.Context().Value(constants.BodyJson)

		bodyAsStruct, ok := body.(schemas.RegistrationBody)

		if !ok {
			helpers.Throw(w, http.StatusInternalServerError, &constants.EmptyString)
			logger.Log.Error(
				errorsApp.Error("userValidationMiddleware: missed field \"" + constants.BodyJson + "\" in ctx"),
			)
			return
		}

		err := validation.ValidateStruct(
			&bodyAsStruct,
			validation.Field(&bodyAsStruct.Name, validation.Required, validation.Length(2, 100)),
			validation.Field(&bodyAsStruct.Email, validation.Required, is.Email),
			validation.Field(&bodyAsStruct.Password, validation.Required, validation.Length(6, 100)),
		)

		if err != nil {
			errorText := err.Error()
			logger.Log.ErrorHttp(req, errorText, http.StatusBadRequest)
			helpers.Throw(w, http.StatusBadRequest, &errorText)
			return
		}

		next.ServeHTTP(w, req)
	})
}
