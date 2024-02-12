package router

import (
	"encoding/json"
	"go-tenancy/api/middleware"
	"go-tenancy/storage/models"
	"go-tenancy/utils/compress"
	"net/http"
)

const useCompress = true

func Authorize(req *http.Request, tenancies *[]*models.Tenancy) bool {
	if len(*tenancies) == 0 {
		return false
	}
	return req.Header.Get("authorization") == (*tenancies)[0].Key
}

func Route(route string, handler http.HandlerFunc) {
	http.Handle(route, middleware.Auth(handler))
}

func Send(data any, res http.ResponseWriter) {
	bytes, err := json.Marshal(data)
	if err != nil {
		http.Error(res, "Failed to compress JSON", http.StatusInternalServerError)
		return
	}

	if useCompress {
		compressedData, err := compress.Compress(bytes)
		if err != nil {
			http.Error(res, "Failed to compress JSON", http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Encoding", "gzip")
		_, err = res.Write(compressedData)
		if err != nil {
			http.Error(res, "Failed to write compressed JSON", http.StatusInternalServerError)
			return
		}
	} else {
		_, err := res.Write(bytes)
		if err != nil {
			http.Error(res, "Failed to write JSON", http.StatusInternalServerError)
			return
		}
	}
}
