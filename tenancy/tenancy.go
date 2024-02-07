package tenancy

import (
	"go-tenancy/gen"
)

type Tenancy struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

func Create() *Tenancy {
	return &Tenancy{
		gen.Id(),
		"",
		gen.Key(),
	}
}

func Name(tenancy *Tenancy, name string) *Tenancy {
	tenancy.Name = name
	return &Tenancy{}
}

func Key(tenancy *Tenancy, key string) *Tenancy {
	tenancy.Key = key
	return &Tenancy{}
}

func Instance(id string, name string, key string) *Tenancy {
	return &Tenancy{
		id,
		name,
		key,
	}
}

func GetId(tenancy *Tenancy) string {
	return tenancy.Id
}

func GetName(tenancy *Tenancy) string {
	return tenancy.Name
}

func GetKey(tenancy *Tenancy) string {
	return tenancy.Key
}
