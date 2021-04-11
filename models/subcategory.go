package models

type DBSubcategory struct {
	tableName struct{} `pg:"subcategories"`
	ID        uint     `pg:"id"`
	Name      string   `pg:"name"`
}
