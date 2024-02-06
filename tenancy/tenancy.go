package tenancy

import (
	"go-tenancy/identifier"
)

type Tenancy struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func Create() *Tenancy {
	return &Tenancy{
		identifier.Get(),
		"",
	}
}

func Name(tenancy *Tenancy, name string) *Tenancy {
	tenancy.Name = name
	return &Tenancy{}
}

func Instance(id string, name string) *Tenancy {
	return &Tenancy{
		id,
		name,
	}
}

func GetId(tenancy *Tenancy) string {
	return tenancy.Id
}

func GetName(tenancy *Tenancy) string {
	return tenancy.Name
}
