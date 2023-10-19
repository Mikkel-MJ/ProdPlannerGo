package repositories

import (
	"context"
	"prodplanner/webapp/database"
	"prodplanner/webapp/models"

	"github.com/jackc/pgx/v5"
)

type TaskRepository struct {
	DB *database.DB
}

func NewTaskRepository(db *database.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (s *TaskRepository) FindAllTasks(tenant_id int) ([]models.Task, error) {
	var tasks []models.Task
	rows, err := s.DB.Pool.Query(context.Background(), `SELECT id, title, tenant_id, note, status, order_id, updated_at, created_at FROM tasks 
    WHERE tenant_id = $1 AND deleted_at IS NULL`, tenant_id)
	if err != nil {
		return tasks, err
	}
	tasks, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.Task])
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (s *TaskRepository) FindTaskById(tenant_id int, id int) (models.Task, error) {
	var task models.Task
	rows, err := s.DB.Pool.Query(context.Background(), `SELECT id, title, tenant_id, note, status, order_id, updated_at, created_at FROM tasks 
    WHERE id = $1 AND tenant_id = $2 AND deleted_at IS NULL`, id, tenant_id)
	if err != nil {
		return task, err
	}
	task, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Task])
	if err != nil {
		return task, err
	}
	return task, nil
}

func (s *TaskRepository) CreateTask(tenant_id int, title string, order_id int64) (models.Task, error) {
	row, err := s.DB.Pool.Query(context.Background(), `INSERT INTO tasks (title, tenant_id, order_id) VALUES ($1, $2, $3)
    RETURNING id, title, tenant_id, note, status, order_id, updated_at, created_at`, title, tenant_id, order_id)
	if err != nil {
		return models.Task{}, err
	}
	task, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.Task])
	if err != nil {
		println(err.Error())
		return models.Task{}, nil
	}

	return task, nil
}

func (s *TaskRepository) CreateTaskRanking(tenant_id, order_id int, task_ids []int) error {
	_, err := s.DB.Pool.Exec(context.Background(), `INSERT INTO task_rankings (tenant_id, order_id, rank) VALUES ($1, $2, $3)`, tenant_id, order_id, task_ids)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func (s *TaskRepository) FindTaskRankingByID(tenant_id, order_id int) (*[]int, error) {
	var ranks *[]int
	err := s.DB.Pool.QueryRow(context.Background(), `SELECT rank FROM task_rankings WHERE tenant_id = $1 AND order_id = $2`, tenant_id, order_id).Scan(&ranks)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return ranks, nil
}

func (s *TaskRepository) DeleteTask(tenant_id int, id int) error {
	_, err := s.DB.Pool.Exec(context.Background(), `UPDATE tasks SET deleted_at = NOW() WHERE id = $1 AND tenant_id = $2`, id, tenant_id)
	if err != nil {
		return err
	}
	return nil
}

// TODO: make this into partial update function
func (s *TaskRepository) UpdateTask(tenant_id int, id int, status int, note string) (models.Task, error) {
	var task models.Task
	row, err := s.DB.Pool.Query(context.Background(), `UPDATE tasks SET  status = $1, note = $2 WHERE id = $3 AND tenant_id = $4
    RETURNING id, title, tenant_id, note, status, order_id, updated_at, created_at`, status, note, id, tenant_id)
	if err != nil {
		return task, err
	}
	task, err = pgx.CollectOneRow(row, pgx.RowToStructByName[models.Task])
	if err != nil {
		return task, err
	}
	return task, nil
}

func (s *TaskRepository) UpdateTaskStatus(tenant_id int, id int, status int) (models.Task, error) {
	var task models.Task
	row, err := s.DB.Pool.Query(context.Background(), `UPDATE tasks SET status = $1 WHERE id = $2 AND tenant_id = $3
    RETURNING id, title, tenant_id, note, status, order_id, updated_at, created_at`, status, id, tenant_id)
	if err != nil {
		return task, err
	}
	task, err = pgx.CollectOneRow(row, pgx.RowToStructByName[models.Task])
	if err != nil {
		return task, err
	}
	return task, nil
}
