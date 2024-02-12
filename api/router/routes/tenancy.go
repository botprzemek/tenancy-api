package routes

import (
	"github.com/dlclark/regexp2"
	"go-tenancy/api/router"
	"go-tenancy/storage/database"
	"go-tenancy/storage/models"
	"net/http"
	"strings"
)

func GetTenancies(res http.ResponseWriter, req *http.Request) {
	var tenancies []*models.Tenancy

	parts := strings.Split(req.URL.Path, "/")

	if parts[1] == "/" || (parts[1] == "" && parts[2] == "") {
		database.Tenancies(&tenancies)
		router.Send(tenancies, res)
		return
	}

	println(parts[2])

	if len(parts[2]) != 8 {
		http.NotFound(res, req)
		return
	}

	{
		regex := regexp2.MustCompile(`^[a-z0-9]{8}$`, 0)
		valid, _ := regex.MatchString(parts[2])

		if !valid {
			http.NotFound(res, req)
			return
		}

		database.TenancyByKey(&tenancies, "id", parts[1])

		if len(tenancies) == 0 {
			http.Error(res, "Tenancy with this Id does not exist", 204)
			return
		}

		if !router.Authorize(req, &tenancies) {
			http.NotFound(res, req)
			return
		}

		var data []*models.Data

		for index := range tenancies {
			data = append(data, tenancies[index].GetData())
		}

		router.Send(data, res)
		return
	}
}
