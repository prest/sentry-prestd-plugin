// nolint
// all plugins must have their package name as `main`
// each plugin is isolated at compile time
package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni/v3"
)

// BUILD:
// go build -o ./lib/middlewares/sentry.so -buildmode=plugin ./lib/src/middlewares/sentry.go
func SentryMiddlewareLoad() negroni.Handler {
	fmt.Println("Plugin chamado")
	return negroni.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request, next http.HandlerFunc) {
		rw.Header().Add("X-Sentry", "Sentry Hello")
		next(rw, rq)
	})
}
