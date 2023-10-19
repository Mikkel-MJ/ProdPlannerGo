package handlers

import (
	"prodplanner/webapp/services"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Route interface {
	fiber.Handler
	Pattern() string
}

type TemplateHandler struct {
	TS *services.TemplateService
}

func NewTemplateHandler(ts *services.TemplateService) *TemplateHandler {
	return &TemplateHandler{TS: ts}
}

func (t *TemplateHandler) GetTaskTemplates(c *fiber.Ctx) error {
	tasks, err := t.TS.FindAllTaskTemplates(1)
	if err != nil {
		return err
	}
	return c.Render("index", fiber.Map{"Tasks": tasks}, "layouts/main")
}

func (t *TemplateHandler) GetOrderTemplateWithTasks(c *fiber.Ctx) error {
	return nil
}

func (t *TemplateHandler) CreateTaskTemplate(c *fiber.Ctx) error {
	value := c.FormValue("title")

	task, err := t.TS.CreateTaskTemplate(1, value)
	if err != nil {
		return err
	}

	return c.Render("taskRow", task)
}

func (t *TemplateHandler) DeleteTaskTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = t.TS.DeleteTaskTemplate(1, i)
	if err != nil {
		return err
	}
	return c.SendString("")
}

func (t *TemplateHandler) GetOrderTemplates(c *fiber.Ctx) error {
	templates, err := t.TS.FindAllOrderTemplates(1)
	if err != nil {
		return err
	}
	tasks, err := t.TS.FindAllTaskTemplates(1)
	if err != nil {
		return err
	}

	return c.Render("order-templates", fiber.Map{"Orders": templates, "Tasks": tasks}, "layouts/main")
}

type createOrderTemplate struct {
	Title string   `form:"title"`
	Tasks []string `form:"ids"`
}

func (t *TemplateHandler) CreateOrderTemplate(c *fiber.Ctx) error {
	order := new(createOrderTemplate)
	if err := c.BodyParser(order); err != nil {
		return err
	}

	println(order.Title + " " + strings.Join(order.Tasks, " "))
	return nil
}
