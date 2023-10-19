package repositories

import (
	"context"
	"prodplanner/webapp/database"
	"prodplanner/webapp/models"

	"github.com/jackc/pgx/v5"
)

type TemplateRepository struct {
	DB *database.DB
}

func NewTemplateService(db *database.DB) *TemplateRepository {
	return &TemplateRepository{DB: db}
}

func (s *TemplateRepository) FindTaskTemplateById(tenant_id int, id int) (models.TaskTemplate, error) {
	var task models.TaskTemplate

	rows, err := s.DB.Pool.Query(context.Background(), `SELECT id, title, tenant_id, created_at FROM task_templates 
    WHERE id = $1 AND tenant_id = $2 AND deleted_at IS NULL`, id, tenant_id)
	if err != nil {
		return task, err
	}
	task, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.TaskTemplate])
	if err != nil {
		return task, err
	}

	return task, nil
}

func (s *TemplateRepository) FindAllTaskTemplates(tenant_id int) ([]models.TaskTemplate, error) {
	var tasks []models.TaskTemplate
	rows, err := s.DB.Pool.Query(context.Background(), `SELECT id, title, tenant_id, created_at FROM task_templates 
    WHERE tenant_id = $1 AND deleted_at IS NULL`, tenant_id)
	if err != nil {
		return tasks, err
	}
	tasks, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.TaskTemplate])
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (s *TemplateRepository) CreateTaskTemplate(tenant_id int, title string) (models.TaskTemplate, error) {
	var task models.TaskTemplate
	row, err := s.DB.Pool.Query(context.Background(), `INSERT INTO task_templates (title, tenant_id) VALUES ($1, $2) RETURNING id, title, tenant_id, created_at`, title, tenant_id)
	if err != nil {
		return task, err
	}
	task, err = pgx.CollectOneRow(row, pgx.RowToStructByName[models.TaskTemplate])
	if err != nil {
		return task, err
	}
	return task, nil
}

func (s *TemplateRepository) DeleteTaskTemplate(tenant_id int, id int) error {
	_, err := s.DB.Pool.Exec(context.Background(), `UPDATE task_templates SET deleted_at = NOW() WHERE id = $1 AND tenant_id = $2`, id, tenant_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TemplateRepository) UpdateTaskTemplateTitle(tenant_id int, id int, title string) (models.TaskTemplate, error) {
	var task models.TaskTemplate
	row, err := s.DB.Pool.Query(context.Background(), `UPDATE task_templates SET title = $1 WHERE id = $2 AND tenant_id = $3 RETURNING id, title, tenant_id, created_at`, title, id, tenant_id)
	if err != nil {
		return task, err
	}
	task, err = pgx.CollectOneRow(row, pgx.RowToStructByName[models.TaskTemplate])
	if err != nil {
		return task, err
	}
	return task, nil
}

func (s *TemplateRepository) FindAllOrderTemplates(tenant_id int) ([]models.OrderTemplate, error) {
	var orderTemplates []models.OrderTemplate
	rows, err := s.DB.Pool.Query(context.Background(), `SELECT id, title, tenant_id FROM order_templates 
    WHERE tenant_id = $1 AND deleted_at IS NULL`, tenant_id)
	if err != nil {
		return orderTemplates, err
	}
	orderTemplates, err = pgx.CollectRows(rows, pgx.RowToStructByName[models.OrderTemplate])
	if err != nil {
		return orderTemplates, err
	}
	return orderTemplates, nil
}

func (s *TemplateRepository) FindOrderTemplateById(tenant_id int, id int) (models.OrderTemplate, error) {
	var orderTemplate models.OrderTemplate
	row, err := s.DB.Pool.Query(context.Background(), `SELECT id, title, tenant_id FROM order_templates 
    WHERE id = $1 AND tenant_id = $2 AND deleted_at IS NULL`, id, tenant_id)
	if err != nil {
		return orderTemplate, err
	}
	orderTemplate, err = pgx.CollectOneRow(row, pgx.RowToStructByName[models.OrderTemplate])
	if err != nil {
		return orderTemplate, err
	}
	return orderTemplate, nil
}

func (s *TemplateRepository) CreateOrderTemplate(tenant_id int, title string) (models.OrderTemplate, error) {
	var orderTemplate models.OrderTemplate
	row, err := s.DB.Pool.Query(context.Background(), `INSERT INTO order_templates (title, tenant_id) VALUES ($1, $2) RETURNING id, title, tenant_id`, title, tenant_id)
	if err != nil {
		return orderTemplate, err
	}
	orderTemplate, err = pgx.CollectOneRow(row, pgx.RowToStructByName[models.OrderTemplate])

	if err != nil {
		return orderTemplate, err
	}
	return orderTemplate, nil
}

func (s *TemplateRepository) DeleteOrderTemplateById(tenant_id int, id int) (models.OrderTemplate, error) {
	var orderTemplate models.OrderTemplate
	row, err := s.DB.Pool.Query(context.Background(), `UPDATE order_templates SET deleted_at = NOW() WHERE id = $1 AND tenant_id = $2 RETURNING id, title, tenant_id`, id, tenant_id)
	if err != nil {
		return orderTemplate, err
	}
	orderTemplate, err = pgx.CollectOneRow(row, pgx.RowToStructByName[models.OrderTemplate])
	if err != nil {
		return orderTemplate, err
	}
	return orderTemplate, nil
}
