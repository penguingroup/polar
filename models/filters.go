package models

type Category struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type City struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
