package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/ManuelP84/calendar/domain/task/models"
)

const (
	DRIVER_NAME = "postgres"
	TASKS_TABLE = "tasks"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(settings *PostgresDbSettings) *PostgresRepository {
	addr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", settings.User, settings.Password, settings.Host, settings.Port, settings.Db)
	db, err := sql.Open(DRIVER_NAME, addr)
	if err != nil {
		panic(err)
	}
	return &PostgresRepository{db}
}

func (repo *PostgresRepository) Close() {
	repo.db.Close()
}

func (repo *PostgresRepository) InsertTask(ctx context.Context, task *models.Task) error {
	query := "INSERT INTO tasks (id, title, description, created_at) values ($1, $2, $3, $4)"
	_, err := repo.db.ExecContext(ctx, query, task.Id, task.Title, task.Description, time.Now())
	return err
}

func (repo *PostgresRepository) UpdateTask(ctx context.Context, task *models.Task) error {
	selectQuery := fmt.Sprintf("SELECT id, title, description, created_at FROM %s WHERE id = $1", TASKS_TABLE)
	updateQuery := "UPDATE tasks SET title = $2, description = $3 WHERE id = $1"
	row := repo.db.QueryRowContext(ctx, selectQuery, task.Id)
	searchedTask := &models.Task{}
	if err := row.Scan(&searchedTask.Id, &searchedTask.Title, &searchedTask.Description, &searchedTask.CreatedAt); err != nil {
		return err
	}
	_, err := repo.db.ExecContext(ctx, updateQuery, task.Id, task.Title, task.Description)
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

func (repo *PostgresRepository) DeleteTaskById(ctx context.Context, id string) error {
	selectQuery := fmt.Sprintf("SELECT id, title, description, created_at FROM %s WHERE id = $1", TASKS_TABLE)
	deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", TASKS_TABLE)
	row := repo.db.QueryRowContext(ctx, selectQuery, id)
	task := &models.Task{}
	if err := row.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedAt); err != nil {
		return err
	}
	_, err := repo.db.ExecContext(ctx, deleteQuery, id)
	return err
}

func (repo *PostgresRepository) SearchTaskById(ctx context.Context, id string) (*models.Task, error) {
	query := fmt.Sprintf("SELECT id, title, description, created_at FROM %s WHERE id = $1", TASKS_TABLE)
	row := repo.db.QueryRowContext(ctx, query, id)
	task := &models.Task{}
	if err := row.Scan(&task.Id, &task.Title, &task.Description, &task.CreatedAt); err != nil {
		return nil, err
	}
	return task, nil
}
