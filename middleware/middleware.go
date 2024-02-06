package middleware

import (
	"log"
	"net/http"
)

func Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Header.Get("x-api-key") != "chujec" {
			http.Error(res, "Failed to authenticate, please provide a valid API Key", http.StatusForbidden)
			return
		}

		log.Printf("[%v] request from http://%v/\n", req.Method, req.RemoteAddr)

		res.Header().Set("Content-Type", "application/json")

		handler(res, req)
	}
}
