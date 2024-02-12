package models

type Tenancy struct {
	Id   string `json:"id"`
	Key  string `json:"key"`
	Data Data   `json:"data"`
}

type Data struct {
	Name string `json:"name"`
}

func (tenancy Tenancy) Create(name string) *Tenancy {
	return &Tenancy{
		gen.Id(),
		gen.Key(),
		Data{
			name,
		},
	}
}

func (data Data) Create(name string) *Data {
	return &Data{
		name,
	}
}

func (tenancy Tenancy) Instance(id string, name string, key string) *Tenancy {
	return &Tenancy{
		id,
		name,
		Data{key},
	}
}

func (tenancy Tenancy) SetName(name string) *Tenancy {
	tenancy.Data.Name = name
	return &Tenancy{}
}

func (tenancy Tenancy) SetKey(key string) *Tenancy {
	tenancy.Key = key
	return &Tenancy{}
}

func (tenancy Tenancy) GetId() string {
	return tenancy.Id
}

func (tenancy Tenancy) GetKey() string {
	return tenancy.Key
}

func (tenancy Tenancy) GetData() *Data {
	return &(tenancy.Data)
}
