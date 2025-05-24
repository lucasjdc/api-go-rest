package main

import (
	"fmt"

	"github.com/lucasjdc/api-go-rest/models"
	"github.com/lucasjdc/api-go-rest/routes"
)

func main() {

	models.Personalidades = []models.Personalidade {
		{Id: 1, Nome: "Cacau", Historia: "Historia 1"},
		{Id: 2, Nome: "Fiona", Historia: "Historia 2"},
		{Id: 3, Nome: "Luna", Historia: "Historia 3"},
		{Id: 4, Nome: "Rex", Historia: "Historia 3"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}