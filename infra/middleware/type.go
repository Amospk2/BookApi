package middleware

import (
	"net/http"
	"strings"
)

func ApplicationTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			if !strings.Contains(r.URL.Path, "/public") {
				w.Header().Set("Content-Type", "application/json")
			}

			next.ServeHTTP(w, r)
		},
	)
}
