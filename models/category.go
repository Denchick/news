package models

type DBCategory struct {
	tableName struct{} `pg:"categories"`
	ID        uint     `pg:"id"`
	Name      string   `pg:"name"`
}
