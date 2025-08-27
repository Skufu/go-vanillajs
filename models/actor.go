package models

type Actor struct {
	ID       int
	Name     string
	LastName string
	ImageURL *string //nullable
}
