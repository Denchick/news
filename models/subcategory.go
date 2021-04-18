package models

// DBSubcategory is the database model for subcategory
type DBSubcategory struct {
	tableName struct{} `pg:"subcategories"`
	ID        uint     `pg:"id"`
	Name      string   `pg:"name"`
}
