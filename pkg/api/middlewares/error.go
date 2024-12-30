package middlewares

import (
	"fmt"
	"net/http"

	"github.com/emmanuerl/vaultly/pkg/internal"
	validation "github.com/go-ozzo/ozzo-validation"
)

var defaultError = "we're currently experiencing system level issues"

func HttpErrorHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				if rvr == http.ErrAbortHandler {
					panic(rvr)
				}

				if err, ok := rvr.(validation.Errors); ok {
					internal.HttpRespond(w, http.StatusUnprocessableEntity, err)
				} else if err, ok := rvr.(*internal.ApiErr); ok {
					internal.HttpRespond(w, err.StatusCode, err)
				} else {
					fmt.Printf("%+v\n", rvr)
					internal.HttpRespond(w, http.StatusInternalServerError, internal.ApiErr{Message: defaultError})
				}
			}
		}()
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
