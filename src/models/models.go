package models

type Magazine struct {
	Id      int     `json:"id"`
	Title   string  `json:"title"`
	Company string  `json:"company"`
	Price   float64 `json:"price"`
	Month   int     `json:"month"`
	Year    int     `json:"year"`
}
