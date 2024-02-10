package router

import (
	"encoding/json"
	"github.com/dlclark/regexp2"
	"go-tenancy/compress"
	"go-tenancy/database"
	"go-tenancy/middleware"
	"go-tenancy/tenancy"
	"net/http"
	"strings"
)

const useCompress = true

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

func Authorize(req *http.Request, tenancies *[]*tenancy.Tenancy) bool {
	if len(*tenancies) == 0 {
		return false
	}
	return req.Header.Get("authorization") == (*tenancies)[0].Key
}

func GetTenancies(res http.ResponseWriter, req *http.Request) {
	var tenancies []*tenancy.Tenancy

	parts := strings.Split(req.URL.Path, "/")

	if parts[0] == "/" || (parts[0] == "" && parts[1] == "") {
		database.Tenancies(&tenancies)
	}

	if len(parts[1]) == 8 {
		regex := regexp2.MustCompile(`^[a-z0-9]{8}$`, 0)
		valid, _ := regex.MatchString(parts[1])

		if !valid {
			http.NotFound(res, req)
			return
		}

		database.TenancyByKey(&tenancies, "id", parts[1])

		if !Authorize(req, &tenancies) {
			http.NotFound(res, req)
			return
		}
	}

	if len(tenancies) == 0 {
		http.Error(res, "Tenancy with this Id does not exist", 204)
		return
	}

	Send(tenancies, res)
}
