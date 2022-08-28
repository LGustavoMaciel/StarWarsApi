package routes

import (
	"StarWarsApi/controllers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)
func Routes(){

	r := chi.NewRouter()
	r.Get("/", controllers.GetPlanets)
	r.Get("/planet/{id}", controllers.GetPlanetById)
	r.Get("/planet/{nome}", controllers.GetPlanetByName)
	r.Post("/planet", controllers.CreatePlanet)
	r.Delete("/planet/{id}", controllers.DeletePlanet)

	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil{
		 panic(err)
	}
}