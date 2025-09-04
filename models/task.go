package models

type Task struct {
	ID          int    `json:"id"` //sera convertido em json qnd for enviado ou recebido
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

const (
	TableName      = "tasks"
	CreateTableSQL = `CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title varchar(100) NOT NULL,
		description TEXT,
		status BOOLEAN NOT NULL DEFAULT false
	)
	`
)
