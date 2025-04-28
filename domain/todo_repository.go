package domain

import "belajar_go/entity"

// TodoRepository is an interface that defines the methods for managing todos
type TodoRepository interface {
	// GetAll retrieves all todos from the repository
	GetAll() ([]entity.Todo, error)
	// GetByID retrieves a todo by its ID
	GetByID(id int) (entity.Todo, error)
	// Create adds a new todo to the repository
	Create(todo entity.Todo) (entity.Todo, error)
	// Update modifies an existing todo in the repository
	Update(todo entity.Todo) (entity.Todo, error)
	// Delete removes a todo from the repository by its ID
	Delete(id int) error
}

