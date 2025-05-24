package models

type Personalidade struct {
	Id       int 	`jsoun:"id"`
	Nome 	 string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade