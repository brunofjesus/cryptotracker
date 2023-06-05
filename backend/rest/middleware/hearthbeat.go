package middleware

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Hearthbeat(next http.Handler) http.Handler {
	return middleware.Heartbeat("/ping")(next)
}
