package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func SampleMiddleware(next http.Handler) http.Handler {
	log.Println("SampleMiddleware")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("SampleMiddleware tick")
		next.ServeHTTP(w, r)
	})
}

func Logging(next http.Handler) http.Handler {
	log.Println("Logging1")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// next.ServeHTTP(w, r)
		//
		// log.Println(r.Method, r.URL.Path, time.Since(start))

		wrappedResponseWriter := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrappedResponseWriter, r)

		log.Println(wrappedResponseWriter.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
