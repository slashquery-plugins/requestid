package requestid

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get("X-Request-ID")
		if rid == "" {
			// Creating UUID Version 1
			uuid1 := uuid.NewV1()
			rid, err := uuid1
			if err != nil {
				log.Println(err)
			}
			r.Header.Set("X-Request-ID", rid)
		}
		next.ServeHTTP(w, r)
	})
}
