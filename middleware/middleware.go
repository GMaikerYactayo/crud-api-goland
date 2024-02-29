package middleware

import (
	"log"
	"net/http"
	"time"
)

func Log(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Petition %q, Method %q", r.URL.Path, r.Method)
		f(w, r)
		log.Printf("Time elapsed duration %q", time.Since(startTime))
	}
}

func Authentication(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		toke := r.Header.Get("Authorization")
		if toke != "token-test" {
			forbidden(w, r)
			return
		}
		f(w, r)
	}
}

func forbidden(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte("Token invalid"))
}
