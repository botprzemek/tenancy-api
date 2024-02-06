package router

import (
	"encoding/json"
	"go-tenancy/compress"
	"go-tenancy/database"
	"go-tenancy/middleware"
	"go-tenancy/tenancy"
	"net/http"
)

const compressing = true

func Route(route string, handler http.HandlerFunc) {
	http.Handle(route, middleware.Auth(handler))
}

func GetTenancies(res http.ResponseWriter, _ *http.Request) {
	var tenancies []*tenancy.Tenancy

	database.Tenancies(&tenancies)

	data, err := json.Marshal(tenancies)
	if err != nil {
		http.Error(res, "Failed to compress JSON", http.StatusInternalServerError)
	}

	if compressing {
		res.Header().Set("Content-Encoding", "gzip")
		data, err = compress.Compress(data)
		if err != nil {
			http.Error(res, "Failed to compress JSON", http.StatusInternalServerError)
			return
		}
	}

	_, err = res.Write(data)
	if err != nil {
		http.Error(res, "Failed to compress JSON", http.StatusInternalServerError)
		return
	}
}
