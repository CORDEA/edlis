package models

type Id struct {
	Id string `json:"id"`
}

func NewId(id string) Id {
	return Id{
		Id: id,
	}
}
