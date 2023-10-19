package repositories

import (
	"context"
	"prodplanner/webapp/database"
	"prodplanner/webapp/models"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type OrderRepository struct {
	DB *database.DB
}

func NewOrderRepository(db *database.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (s *OrderRepository) FindOrderById(tenant_id int, id int) (models.Order, error) {
	var order models.Order
	rows, err := s.DB.Pool.Query(context.Background(), `SELECT id, title, tenant_id, note, status, order_nr, deadline_at, updated_at, created_at FROM orders 
    WHERE id = $1 AND tenant_id = $2 AND deleted_at IS NULL`, id, tenant_id)
	if err != nil {
		return order, err
	}
	order, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Order])
	if err != nil {
		return order, err
	}
	return order, nil
}

func (s *OrderRepository) FindAllOrders(tenant_id int) ([]*models.Order, error) {
	var orders []*models.Order
	err := pgxscan.Select(context.Background(), s.DB.Pool, &orders, `SELECT id, title, tenant_id, note, status, order_nr, deadline_at, updated_at, created_at FROM orders 
    WHERE tenant_id = $1 AND deleted_at IS NULL`, tenant_id)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderRepository) FindAllOrdersWithTasks(tenant_id int) ([]models.OrderTasks, error) {
	rows, err := s.DB.Pool.Query(context.Background(), `
select o.title,o.order_nr,o.id,o.status, t.title as task_title, t.status as task_status, t.id as task_id
from orders o
JOIN public.tasks t on o.id = t.order_id
WHERE o.tenant_id = $1`, tenant_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	orderMap := make(map[int]models.OrderTasks)
	for rows.Next() {
		var title, order_nr, task_title string
		var id, task_id, status, task_status int
		err = rows.Scan(&title, &order_nr, &id, &status, &task_title, &task_status, &task_id)
		if err != nil {
			return nil, err
		}
		orderTask, ok := orderMap[id]
		if !ok {
			orderTask = models.OrderTasks{
				Title:        title,
				Order_number: order_nr,
				ID:           int64(id),
				Status:       models.Status(status),
			}
		}
		task := models.SimpleTask{Title: task_title, ID: int64(task_id), Status: models.Status(task_status)}
		orderTask.Tasks = append(orderTask.Tasks, task)

		orderMap[id] = orderTask
	}
	if err := rows.Err(); err != nil {
		// Handle any rows.Err() error, if necessary.
		return nil, err
	}

	var orders []models.OrderTasks
	for _, v := range orderMap {
		orders = append(orders, v)
	}

	return orders, nil
}

func (s *OrderRepository) CreateOrder(tenant_id int, title string, note string, deadline_at time.Time, order_nr string) (models.Order, error) {
	var order models.Order
	row, err := s.DB.Pool.Query(context.Background(), `INSERT INTO orders (title, tenant_id, note, deadline_at, order_nr) VALUES ($1, $2, $3, $4, $5)
    RETURNING id, title, tenant_id, note, status, order_nr, deadline_at, updated_at, created_at`, title, tenant_id, note, deadline_at, order_nr)
	if err != nil {
		return order, err
	}
	order, err = pgx.CollectOneRow(row, pgx.RowToStructByName[models.Order])
	if err != nil {
		return order, err
	}
	return order, nil
}

func (s *OrderRepository) DeleteOrder(tenant_id int, id int) (models.Order, error) {
	var order models.Order
	row, err := s.DB.Pool.Query(context.Background(), `UPDATE orders SET deleted_at = NOW() WHERE id = $1 AND tenant_id = $2
    RETURNING id, title, tenant_id, note, status, order_nr, deadline_at, updated_at, created_at`, id, tenant_id)
	if err != nil {
		return order, err
	}
	order, err = pgx.CollectOneRow(row, pgx.RowToStructByName[models.Order])
	if err != nil {
		return order, err
	}
	return order, nil
}

func (s *OrderRepository) UpdateOrderTitle(tenant_id int, id int, title string) (models.Order, error) {
	var order models.Order
	row, err := s.DB.Pool.Query(context.Background(), `UPDATE orders SET title = $1 WHERE id = $2 AND tenant_id = $3
    RETURNING id, title, tenant_id, note, status, order_nr, deadline_at, updated_at, created_at`, title, id, tenant_id)
	if err != nil {
		return order, err
	}
	order, err = pgx.CollectOneRow(row, pgx.RowToStructByName[models.Order])
	if err != nil {
		return order, err
	}
	return order, nil
}
