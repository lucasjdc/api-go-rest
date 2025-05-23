package main

import (
	"fmt"

	"github.com/lucasjdc/api-go-rest/models"
	"github.com/lucasjdc/api-go-rest/routes"
)

func main() {

	models.Personalidades = []models.Personalidade {
		{Nome: "Cacau", Historia: "Historia 1"},
		{Nome: "Fiona", Historia: "Historia 2"},
		{Nome: "Luna", Historia: "Historia 3"},
		{Nome: "Rex", Historia: "Historia 3"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}