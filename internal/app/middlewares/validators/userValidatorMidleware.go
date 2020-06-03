package validators

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/stetsd/monk-api/internal/app/constants"
	"github.com/stetsd/monk-api/internal/app/schemas"
	"net/http"
)

func Registration(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body := req.Context().Value(constants.BodyJson)

		if body == nil {
			handleErrorNoBody("userValidationMiddleware: missed field \""+constants.BodyJson+"\" in ctx", w)
			return
		}

		bodyAsStruct, ok := body.(schemas.RegistrationBody)

		if !ok {
			handleErrorNoBody("userValidationMiddleware: json parse error", w)
			return
		}

		err := validation.ValidateStruct(
			&bodyAsStruct,
			validation.Field(&bodyAsStruct.Name, validation.Required, validation.Length(2, 100)),
			validation.Field(&bodyAsStruct.Email, validation.Required, is.Email),
			validation.Field(&bodyAsStruct.Password, validation.Required, validation.Length(6, 100)),
		)

		if err != nil {
			handleErrorValidation(err, w, req)
			return
		}

		next.ServeHTTP(w, req)
	})
}
