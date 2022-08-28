package models

import "gopkg.in/validator.v2"


type Planet struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Clima string `json:"clima"`
	Terreno string `json:"terreno"`
}


func ValidatePlanet(planet *Planet) error {

	err := validator.Validate(planet)

	if err != nil {
		return err
	}

	return nil 
}