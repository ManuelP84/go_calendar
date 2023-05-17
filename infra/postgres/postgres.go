package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/ManuelP84/calendar/domain/task/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) Close() {
	repo.db.Close()
}

func (repo *PostgresRepository) InsertTask(ctx context.Context, task *models.Task) error {
	insert := "INSERT INTO tasks (id, title, description) values ($1, $2, $3)"
	_, err := repo.db.ExecContext(ctx, insert, task.Id, task.Title, task.Description)
	return err
}

func (repo *PostgresRepository) GetTasks(ctx context.Context) ([]*models.Task, error) {
	query := "SELECT id, title, description, created_at FROM tasks"
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := []*models.Task{}

	for rows.Next() {
		task := &models.Task{}
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
