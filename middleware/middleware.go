package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func Auth(handler http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Header.Get("x-api-key") != os.Getenv("API_KEY") {
			log.Printf("[%v] request rejected from %v\n", req.Method, strings.Split(req.RemoteAddr, ":")[0])

			http.Error(res, "Failed to authenticate, please provide a valid API Key", http.StatusForbidden)

			return
		}

		log.Printf("[%v] request successed from %v/\n", req.Method, strings.Split(req.RemoteAddr, ":")[0])

		res.Header().Set("Content-Type", "application/json")

		handler(res, req)
	}
}
