package datamodels

import (
	"database/sql"
	"time"
)

// Snippet hold individual snippet from database
type Snippet struct {
	ID           int64
	Title        string
	Content      string
	CreationTime time.Time
	ExpriyTime   time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (snippetModel *SnippetModel) InsertASnippet(
	title string,
	content string,
	expiressAt int,
) (int, error) {
	query := `INSERT INTO snippets (title, content, created, expires)
  VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, sqlError := snippetModel.DB.Exec(query, title, content, expiressAt)
	if sqlError != nil {
		return 0, sqlError
	}
	id, sqlError := result.LastInsertId()
	if sqlError != nil {
		return 0, nil
	}
	return int(id), nil
}
