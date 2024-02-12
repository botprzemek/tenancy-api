package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"
)

const useKey = true

func Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if useKey && req.Header.Get("x-api-key") != os.Getenv("API_KEY") {
			log.Printf("[%v] request rejected from %v\n", req.Method, strings.Split(req.RemoteAddr, ":")[0])

			http.Error(res, "Failed to authenticate, please provide a valid API Key", http.StatusForbidden)
			return
		}

		log.Printf("[%v] request succeed from %v\n", req.Method, strings.Split(req.RemoteAddr, ":")[0]+req.RequestURI)

		res.Header().Set("Content-Type", "application/json")
		handler(res, req)
	}
}
