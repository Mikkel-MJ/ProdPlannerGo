package services

import (
	"fmt"
	"prodplanner/webapp/models"
	"prodplanner/webapp/repositories"
	"time"
)

type OrderService struct {
	OR *repositories.OrderRepository
	TR *repositories.TaskRepository
}

func NewOrderService(or *repositories.OrderRepository, tr *repositories.TaskRepository) *OrderService {
	return &OrderService{OR: or, TR: tr}
}

func (s *OrderService) FindOrderById(tenant_id int, id int) (models.Order, error) {
	order, err := s.OR.FindOrderById(tenant_id, id)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (s *OrderService) FindOrders(tenant_id int) ([]*models.Order, error) {
	orders, err := s.OR.FindAllOrders(tenant_id)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// TODO: This is a mess, clean it up and make it more performant
func (s *OrderService) FindOrdersWithTasks(tenant_id int) ([]models.OrderTasks, error) {
	orders, err := s.OR.FindAllOrdersWithTasks(tenant_id)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(orders); i++ {
		ranks, _ := s.TR.FindTaskRankingByID(1, int(orders[i].ID))
		idToTaskMap := make(map[int]models.SimpleTask)
		for _, task := range orders[i].Tasks {
			idToTaskMap[int(task.ID)] = task
		}
		var sortedTasks []models.SimpleTask
		for _, id := range *ranks {
			sortedTasks = append(sortedTasks, idToTaskMap[id])
		}
		orders[i].Tasks = sortedTasks
	}

	return orders, nil
}

func (s *OrderService) CreateOrder(tenant_id int, req models.CreateOrderRequest) (models.Order, error) {
	date, err := time.Parse("02-01-2006", req.Deadline)
	if err != nil {
		return models.Order{}, err
	}
	order, err := s.OR.CreateOrder(tenant_id, req.Title, req.Note, date, req.OrderNum)
	if err != nil {
		return models.Order{}, err
	}
	ranks := make([]int, 0, len(req.Tasks))
	for _, task := range req.Tasks {
		t, err := s.TR.CreateTask(tenant_id, task, order.ID)
		if err != nil {
			return models.Order{}, err
		}
		println(t.ID)
		ranks = append(ranks, int(t.ID))
	}
	fmt.Println(ranks)
	s.TR.CreateTaskRanking(tenant_id, int(order.ID), ranks)
	return order, nil
}
