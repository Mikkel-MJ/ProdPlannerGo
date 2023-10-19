package models

import "time"

type TaskTemplate struct {
	CreatedAt time.Time `db:"created_at"`
	Title     string    `db:"title"`
	ID        int64     `db:"id"`
	TenantID  int64     `db:"tenant_id"`
}

type OrderTemplate struct {
	Title    string `db:"title"`
	ID       int64  `db:"id"`
	TenantID int64  `db:"tenant_id"`
}

type OrderTemplateWithTasks struct {
	Title         string
	Description   string
	TaskTemplates []TaskTemplate
	Id            int64
	TenantID      int64
}

// Order models
type Status int // 0 - created, 1 - in progress, 2 - done

const (
	NotStarted Status = iota
	InProgress Status = iota
	Completed  Status = iota
	HasIssues  Status = iota
	Blocked    Status = iota
)

var statusNames = [...]string{
	"NotStarted",
	"InProgress",
	"Completed",
	"HasIssues",
	"Blocked",
}

func GetStatusName(status Status) string {
	if status >= 0 && int(status) < len(statusNames) {
		return statusNames[status]
	}
	return "Unknown"
}

type Order struct {
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	DeadlineAT   time.Time `db:"deadline_at"`
	Title        string    `db:"title"`
	Order_number string    `db:"order_nr"`
	Note         string    `db:"note"`
	ID           int64     `db:"id"`
	TenantID     int64     `db:"tenant_id"`
	Status       Status    `db:"status"`
}
type OrderTasks struct {
	Title        string
	Order_number string
	Tasks        []SimpleTask
	ID           int64
	Status       Status
}

type SimpleTask struct {
	Title  string
	ID     int64
	Status Status
}
type OrderWithTasks struct {
	Title        string `db:"title"`
	Order_number string `db:"order_nr"`
	TaskTitle    string `db:"task_title"`
	ID           int64  `db:"id"`
	Status       Status `db:"status"`
	TaskID       int64  `db:"task_id"`
	TaskStatus   Status `db:"task_status"`
}

// Task models
type Task struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Title     string    `db:"title"`
	Note      string    `db:"note"`
	ID        int64     `db:"id"`
	OrderID   int64     `db:"order_id"`
	TenantID  int64     `db:"tenant_id"`
	Status    Status    `db:"status"`
}

type CreateOrderRequest struct {
	Title    string   `form:"title"`
	OrderNum string   `form:"order_nr"`
	Note     string   `form:"note"`
	Deadline string   `form:"deadline"`
	Tasks    []string `form:"titles"`
}

type UpdateTaskRequest struct {
	Note   string `form:"note"`
	Status int    `form:"status"`
}
type CreateTask struct {
	Title string
	Rank  string
}
