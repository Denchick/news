package models

// DBCategory is the database model for Category
type DBCategory struct {
	tableName struct{} `pg:"categories"`
	ID        uint     `pg:"id"`
	Name      string   `pg:"name"`
}
