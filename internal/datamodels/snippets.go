package datamodels

import (
	"database/sql"
	"errors"
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

// Insert/Create a new Snippet
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

// Returns the snippet for the given snippetId
func (snippetModel *SnippetModel) GetSnippetById(snippetId int) (*Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets where expires > UTC_TIMESTAMP AND id = ?`
	row := snippetModel.DB.QueryRow(query, snippetId)
	result := &Snippet{}
	// copy value of each field to the new snippetObject
	err := row.Scan(&result.ID, &result.Title, &result.Content, &result.CreationTime, &result.ExpriyTime)
	if err != nil {
		// if the snippetId has no rows
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			// any other sql relatedErr
			return nil, err
		}
	}
	return result, nil
}

// Returns all the lastest snippets of size 10
func (snippetModel *SnippetModel) GetLatestiSnippets() ([]*Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snippets where expires > UTC_TIMESTAMP ORDER BY id DESC LIMIT 10`
	rows, sqlError := snippetModel.DB.Query(query)
	if sqlError != nil {
		return nil, sqlError
	}
	// close the db connection
	defer rows.Close()

	// create a slice to hold all the snippets
	snippetList := []*Snippet{}

	// iteration over the resultSet of the query
	for rows.Next() {
		result := &Snippet{}

		sqlError = rows.Scan(&result.ID, &result.Title, &result.Content, &result.CreationTime, &result.ExpriyTime)
		if sqlError != nil {
			return nil, sqlError
		}

		snippetList = append(snippetList, result)
	}
	// to retrive any err faced during iteration of the result set
	if sqlError = rows.Err(); sqlError != nil {
		return nil, sqlError
	}

	return snippetList, nil
}
