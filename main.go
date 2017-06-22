package requestid

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get("X-Request-ID")
		if rid == "" {
			// Creating UUID Version 1
			r.Header.Set("X-Request-ID", fmt.Sprintf("%s", uuid.NewV1()))
		}
		next.ServeHTTP(w, r)
	})
}
