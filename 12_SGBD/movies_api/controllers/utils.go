package controllers

import (
	"errors"

	"github.com/tomasda7/movies_api/services"
)

func MovieExists(id int) error {
	movies, err := services.GetAllMovies()

	if err != nil {
		return err
	}

	for _, m := range movies {
		if m.ID == id {
			return nil
		}
	}

	return errors.New("the movie with the ID provided doesn't exist")
}
