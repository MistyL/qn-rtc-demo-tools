package models

type Sallers struct {
	Id       int64
	Name     string `json:"name"`
	Password string `json:"password"`
	Room     string `json:"room"`
}
