package services

import (
	"prodplanner/webapp/models"
	"prodplanner/webapp/repositories"
)

type TemplateService struct {
	TR *repositories.TemplateRepository
}

func NewTemplateService(tb *repositories.TemplateRepository) *TemplateService {
	return &TemplateService{TR: tb}
}

func (s *TemplateService) FindTaskTemplateById(tenant_id int, id int) (models.TaskTemplate, error) {
	task, err := s.TR.FindTaskTemplateById(tenant_id, id)
	if err != nil {
		return models.TaskTemplate{}, err
	}
	return task, nil
}

func (s *TemplateService) FindAllTaskTemplates(tenant_id int) ([]models.TaskTemplate, error) {
	var tasks []models.TaskTemplate
	tasks, err := s.TR.FindAllTaskTemplates(tenant_id)
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (s *TemplateService) CreateTaskTemplate(tenant_id int, title string) (models.TaskTemplate, error) {
	var task models.TaskTemplate
	task, err := s.TR.CreateTaskTemplate(tenant_id, title)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (s *TemplateService) DeleteTaskTemplate(tenant_id int, id int) error {
	err := s.TR.DeleteTaskTemplate(tenant_id, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *TemplateService) UpdateTaskTemplateTitle(tenant_id int, id int, title string) (models.TaskTemplate, error) {
	var task models.TaskTemplate
	task, err := s.TR.UpdateTaskTemplateTitle(tenant_id, id, title)
	if err != nil {
		return task, err
	}
	return task, nil
}

func (s *TemplateService) FindAllOrderTemplates(tenant_id int) ([]models.OrderTemplate, error) {
	var orderTemplates []models.OrderTemplate
	orderTemplates, err := s.TR.FindAllOrderTemplates(tenant_id)
	if err != nil {
		return orderTemplates, err
	}
	return orderTemplates, nil
}

func (s *TemplateService) FindOrderTemplateById(tenant_id int, id int) (models.OrderTemplate, error) {
	var orderTemplate models.OrderTemplate
	orderTemplate, err := s.TR.FindOrderTemplateById(tenant_id, id)
	if err != nil {
		return orderTemplate, err
	}
	return orderTemplate, nil
}

func (s *TemplateService) CreateOrderTemplate(tenant_id int, title string) (models.OrderTemplate, error) {
	var orderTemplate models.OrderTemplate
	orderTemplate, err := s.TR.CreateOrderTemplate(tenant_id, title)
	if err != nil {
		return orderTemplate, err
	}
	return orderTemplate, nil
}

func (s *TemplateService) DeleteOrderTemplate(tenant_id int, id int) (models.OrderTemplate, error) {
	var orderTemplate models.OrderTemplate
	orderTemplate, err := s.TR.DeleteOrderTemplateById(tenant_id, id)
	if err != nil {
		return orderTemplate, err
	}
	return orderTemplate, nil
}
