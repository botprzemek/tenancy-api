package tenancy

import (
	"go-tenancy/gen"
)

type Tenancy struct {
	Id   string `json:"id"`
	Key  string `json:"key"`
	Data Data   `json:"data"`
}

type Data struct {
	Name string `json:"name"`
}

func Create() *Tenancy {
	return &Tenancy{
		gen.Id(),
		gen.Key(),
		Data{""},
	}
}

func SetName(tenancy *Tenancy, name string) *Tenancy {
	tenancy.Data.Name = name
	return &Tenancy{}
}

func SetKey(tenancy *Tenancy, key string) *Tenancy {
	tenancy.Key = key
	return &Tenancy{}
}

func Instance(id string, name string, key string) *Tenancy {
	return &Tenancy{
		id,
		name,
		Data{key},
	}
}

func GetId(tenancy *Tenancy) string {
	return tenancy.Id
}

func GetKey(tenancy *Tenancy) string {
	return tenancy.Key
}

func GetData(tenancy *Tenancy) *Data {
	return &tenancy.Data
}
