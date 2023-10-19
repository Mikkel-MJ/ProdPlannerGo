package main

import (
	"context"
	"log"
	"prodplanner/webapp/database"
	"prodplanner/webapp/handlers"
	"prodplanner/webapp/models"
	"prodplanner/webapp/repositories"
	"prodplanner/webapp/services"
	"prodplanner/webapp/utils"
	"text/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/fx"
)

func InitApp(lc fx.Lifecycle, th *handlers.TemplateHandler, oh *handlers.OrderHandler) *fiber.App {
	funcmap := template.FuncMap{"getStatusString": models.GetStatusName}
	engine := html.New("./views", ".html")
	engine.AddFuncMap(funcmap)
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	r, _ := utils.Rank("aaabbcg", "aabbddg")
	println("rank: " + r)
	app.Get("/hi", th.GetTaskTemplates)
	app.Post("/hi", th.CreateTaskTemplate)
	app.Delete("/tasks/:id", th.DeleteTaskTemplate)
	app.Get("/templates/order", th.GetOrderTemplates)
	app.Post("/templates/order", th.CreateOrderTemplate)
	app.Get("/order/create", oh.CreateOrderPage)
	app.Post("/order/create", oh.CreateOrder)
	app.Get("/order", oh.OrderPage)
	app.Get("/task/edit/:id", oh.GetEditTaskModal)
	app.Post("/task/edit/:id", oh.EditTask)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Fatal(app.Listen("127.0.0.1:3000"))
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}

func main() {
	fx.New(fx.Provide(
		handlers.NewTemplateHandler, handlers.NewOrderHandler,
		database.NewDatabasePool,
		services.NewTemplateService, services.NewOrderService,
		repositories.NewTaskRepository, repositories.NewTemplateService, repositories.NewOrderRepository,
	), fx.Invoke(InitApp)).Run()
}
