package models

type Movie struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Year     int     `json:"year"`
	Genre    string  `json:"genre"`
	Acquired string  `json:"acquired"`
	Stock    int     `json:"stock"`
	Price    float64 `json:"price"`
}
