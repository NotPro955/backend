package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey int

const (
	requestIDkey ctxKey = iota
)

func RequestId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("X-Request-Id")
		if id == "" {
			id = uuid.NewString()
		}

		w.Header().Add("X-Request-Id", id)
		ctx := context.WithValue(r.Context(), requestIDkey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
