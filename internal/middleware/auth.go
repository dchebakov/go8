package middleware

import (
	"context"
	"fmt"
	"net/http"

	"eight/pkg/masuk"
)

func Authenticate(authenticator masuk.Authenticator, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := authenticator.Authenticate(r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		_ = fmt.Sprintf("user %s Authenticated\n", user.Username())
		f(w, r.WithContext(context.WithValue(r.Context(), "username", user.Username())))
	}
}
