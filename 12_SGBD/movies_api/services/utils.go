package services

import "github.com/tomasda7/movies_api/models"

type RowScanner interface {
	Scan(dest ...interface{}) error
}

func scanMovies(rows RowScanner) (*models.Movie, error) {
	var movie models.Movie

	err := rows.Scan(&movie.ID, &movie.Name, &movie.Year, &movie.Genre, &movie.Acquired, &movie.Stock, &movie.Price)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}
