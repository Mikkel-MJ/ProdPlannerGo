package handlers

import (
	"prodplanner/webapp/models"
	"prodplanner/webapp/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	OS *services.OrderService
	TS *services.TemplateService
}

func NewOrderHandler(os *services.OrderService, ts *services.TemplateService) *OrderHandler {
	return &OrderHandler{OS: os, TS: ts}
}

func (o *OrderHandler) OrderPage(c *fiber.Ctx) error {
	orders, err := o.OS.FindOrdersWithTasks(1)
	if err != nil {
		return err
	}
	return c.Render("order-view", fiber.Map{"Orders": orders}, "layouts/main")
}

func (o *OrderHandler) CreateOrderPage(c *fiber.Ctx) error {
	tasks, err := o.TS.FindAllTaskTemplates(1)
	if err != nil {
		return err
	}
	return c.Render("create-order", fiber.Map{"Tasks": tasks}, "layouts/main")
}

func (o *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	order := new(models.CreateOrderRequest)
	if err := c.BodyParser(order); err != nil {
		return err
	}
	_, err := o.OS.CreateOrder(1, *order)
	println(err)
	return nil
}

func (o *OrderHandler) GetEditTaskModal(c *fiber.Ctx) error {
	id := c.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	task, err := o.OS.TR.FindTaskById(1, i)
	if err != nil {
		return err
	}
	return c.Render("partials/task-edit-modal", fiber.Map{"Task": task})
}

func (o *OrderHandler) EditTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	request := new(models.UpdateTaskRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}
	_, err = o.OS.TR.UpdateTask(1, id, request.Status, request.Note)
	if err != nil {
		return err
	}
	return nil
}
