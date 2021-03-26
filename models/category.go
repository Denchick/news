package models

type DBCategory struct {
	tableName struct{} `pg:"categories"`
	ID        int      `pg:"id"`
	Name      string   `pg:"name"`
}
