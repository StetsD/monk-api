package validators

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/stetsd/monk-api/internal/app/constants"
	"github.com/stetsd/monk-api/internal/app/schemas"
	"net/http"
)

func EventCreate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body := req.Context().Value(constants.BodyJson)

		if body == nil {
			handleErrorNoBody("eventValidationMiddleware: missed field \""+constants.BodyJson+"\" in ctx", w)
			return
		}

		bodyAsStruct, ok := body.(schemas.EventBody)

		fmt.Println(body)

		if !ok {
			handleErrorNoBody("eventValidationMiddleware: json parse error", w)
			return
		}

		err := validation.ValidateStruct(
			&bodyAsStruct,
			validation.Field(&bodyAsStruct.Title, validation.Required),
			validation.Field(&bodyAsStruct.Description, validation.Required),
			validation.Field(&bodyAsStruct.UserId, validation.Required),
			validation.Field(&bodyAsStruct.DateEnd, validation.Required),
			validation.Field(&bodyAsStruct.DateStart, validation.Required),
			validation.Field(&bodyAsStruct.Email, validation.Required, is.Email),
		)

		if err != nil {
			handleErrorValidation(err, w, req)
			return
		}

		next.ServeHTTP(w, req)
	})
}
