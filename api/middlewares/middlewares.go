package middlewares

import (
	"errors"
	"main/api/auth"
	"main/api/responses"
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc, enforceAuth bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if enforceAuth {
			err := auth.TokenValid(r)
			if err != nil {
				responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
				return
			}
		}
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
