package routes

import (
	"log"
	"net/http"

	"github.com/lucasjdc/api-go-rest/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}