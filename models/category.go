package models

type Category struct {
	ID int `pg:"id"`
	Name string `pg:"name"`
}