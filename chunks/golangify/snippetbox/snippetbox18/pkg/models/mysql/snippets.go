package mysql

import (
  "database/sql"
  "golangify.com/snippetbox/snippetbox18/pkg/models" // Уточнить путь.
)

// Определяем тип, который обертывает пул подключения sql.DB
type SnippetModel struct {
  DB *sql.DB
}

// Метод для создания новой заметки в базе данных
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
  return 0, nil
}

// Метод для возврата данных заметки по ее ID
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
  return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
  return nil, nil
}
