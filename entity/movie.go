package entity

type Movie struct {
	ID     string
	Name   string
	Link   string
	Rating int
	MovieCreator
}

type MovieCreator struct {
	CreatorName string `gorm:"column:name"`
	Company    string
}

type MovieRequest struct {
	ID     string `validate:"required"`
	Name   string `validate:"required"`
	Link   string
	Rating int
	Type   string
}
