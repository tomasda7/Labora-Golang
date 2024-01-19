package services

import (
	"errors"

	"github.com/tomasda7/movies_api/models"
)

func GetAllMovies() ([]models.Movie, error) {
	rows, err := DB.Query(`SELECT id, name, year, genre, acquired, stock, price FROM movies;`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movies := []models.Movie{}

	for rows.Next() {
		movie, err := scanMovies(rows)

		if err != nil {
			return nil, err
		}
		movies = append(movies, *movie)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func CreateMovie(newMovie models.Movie) (string, error) {

	// fields validation
	if newMovie.Name == "" || newMovie.Genre == "" || newMovie.Year == 0 || newMovie.Price == 0 {
		return "", errors.New("all fields except 'acquired' and 'stock' should be filled")
	}

	// preparing the SQL statement
	stmt, err := DB.Prepare("INSERT INTO movies (name, year, genre, acquired, stock, price) VALUES ($1, $2, $3, $4, $5, $6) RETURNING name")

	if err != nil {
		return "", err
	}
	defer stmt.Close()

	// executing the SQL statement and retrieving the movie name
	var newMovieName string

	err = stmt.QueryRow(newMovie.Name, newMovie.Year, newMovie.Genre, newMovie.Acquired, newMovie.Stock, newMovie.Price).Scan(&newMovieName)

	if err != nil {
		return "", err
	}

	return newMovieName, nil
}

func UpdateMovie(id int, updMovie models.Movie) (int, error) {

	if updMovie.Name == "" || updMovie.Genre == "" || updMovie.Year == 0 || updMovie.Price == 0 {
		return 0, errors.New("all fields except 'acquired' and 'stock' should be filled")
	}

	stmt, err := DB.Prepare("UPDATE movies SET name=$1, year=$2, genre=$3, acquired=$4, stock=$5, price=$6 WHERE id=$7 RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var updMovieID int

	err = stmt.QueryRow(updMovie.Name, updMovie.Year, updMovie.Genre, updMovie.Acquired, updMovie.Stock, updMovie.Price, id).Scan(&updMovieID)

	if err != nil {
		return 0, err
	}

	return updMovieID, nil
}

func DeleteMovie(id int) (string, error){
	stmt, err := DB.Prepare("DELETE FROM movies WHERE id=$1 RETURNING name")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var dltdMovieName string

	err = stmt.QueryRow(id).Scan(&dltdMovieName)

	if err != nil {
		return "", err
	}

	return dltdMovieName, nil
}
