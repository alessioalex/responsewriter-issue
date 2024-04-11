package middleware

import (
	"fmt"
	"net/http"
)

type Middleware func(h http.Handler) http.Handler

func CreateStack(stack ...Middleware) Middleware {
	fmt.Printf("stack length: %+v\n", len(stack))

	return func(next http.Handler) http.Handler {
		for i := len(stack) - 1; i >= 0; i-- {
			currentMiddleware := stack[i]
			next = currentMiddleware(next)
		}

		return next
	}
}
