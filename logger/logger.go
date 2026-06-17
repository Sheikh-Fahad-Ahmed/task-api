package logger

import (
	"log"
	"net/http"
	"time"
)

func Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		f(w, r)
		log.Printf("[%s] %s | %v", r.Method, r.URL.Path, time.Since(start))
	}
}
