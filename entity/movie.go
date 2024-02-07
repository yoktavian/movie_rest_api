package entity

type Movie struct {
	ID     string
	Name   string
	Link   string
	Rating int
}

type MovieRequest struct {
	ID     string `validate:"required"`
	Name   string
	Link   string
	Rating int
}
