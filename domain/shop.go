package domain

type Shop struct {
	GormModel
	Name string `json:"name"`
	Domain string `json:"domain"`
}
