package repository

import (
	"github.com/DevAthhh/DoZen/internal/models"
	"gorm.io/gorm"
)

// TODO: исправить UpdateTaskStatusByID и сделать так чтобы он принимал только значение статуса, и сам выставлял статус по ключу

type TaskRepository struct {
	db *gorm.DB
}

func (t *TaskRepository) CreateTask(groupID int, title string) error {
	task := models.Task{
		Title:   title,
		Status:  "not marked up",
		Done:    false,
		GroupID: uint(groupID),
	}

	return t.db.Create(&task).Error
}

func (t *TaskRepository) GetTasksByGroupID(id string) (*[]models.Task, error) {
	var tasks []models.Task
	if err := t.db.Where("group_id = ?", id).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return &tasks, nil
}

func (t *TaskRepository) UpdateTaskStatusByID(id, status string) error {
	var task models.Task
	if err := t.db.First(&task, "id = ?", id).Error; err != nil {
		return err
	}

	task.Status = status
	if status == "done" {
		task.Done = true
	}

	return t.db.Save(&task).Error
}

func (t *TaskRepository) DeleteTaskByID(id string) error {
	var task models.Task
	return t.db.Where("id = ?", id).Delete(&task).Error
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}
