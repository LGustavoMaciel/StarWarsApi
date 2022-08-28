package controllers

import (
	"StarWarsApi/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

//CreatePlanet
func CreatePlanet(w http.ResponseWriter, r *http.Request){

	var planet models.Planet

	err := json.NewDecoder(r.Body).Decode(&planet)
	if err != nil {
		log.Printf("Error to send body of request%v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			// "Error":   false,
			"Message": fmt.Sprintf("Dados inseridos com sucesso!"),
		}
	}
	w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
}


//List all planets 
func GetPlanets(w http.ResponseWriter, r *http.Request) {
	

	resp, err := http.Get("https://swapi.dev/api/planets/")
		if err != nil {
			panic(err)
		}
	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, resp.Body) 
	err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


//List by Id
func GetPlanetById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listagem por ID"))
}

//List by name
func GetPlanetByName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listagem por nome"))
}

func DeletePlanet(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("DeletePlanet"))
}