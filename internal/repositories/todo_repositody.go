package repositories

import (
	"tracker/internal/models"
)

type TodoRepository interface {
	List(num int64) ([]*models.Todo, error)
	GetByID(id int64) (*models.Todo, error)
	Create(todo *models.Todo) (int64, error)
	Update(todo *models.Todo) (*models.Todo, error)
	Delete(id int64) (bool, error)
}
