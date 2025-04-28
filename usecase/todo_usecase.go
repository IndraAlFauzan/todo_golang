package usecase

import (
	"belajar_go/apperror"
	"belajar_go/domain"
	"belajar_go/entity"
)

type TodoUsecase interface {
	GetAll() ([]entity.Todo, error)
	GetByID(id int) (entity.Todo, error)
	Create(todo entity.Todo) (entity.Todo, error)
	Update(todo entity.Todo) (entity.Todo, error)
	Delete(id int) error
}

type todoUsecaseImpl struct {
	todoRepo domain.TodoRepository
}

func NewTodoUsecase(todoUsecase domain.TodoRepository) TodoUsecase {
	return &todoUsecaseImpl{
		todoRepo: todoUsecase,
	}
}

// Create implements TodoUsecase.
func (t *todoUsecaseImpl) Create(todo entity.Todo) (entity.Todo, error) {
	if todo.Title == "" {
		return entity.Todo{}, apperror.ValidationError("Title ")
	}
	if todo.CreatedAt == "" {
		return entity.Todo{}, apperror.ValidationError("CreatedAt ")
	}

	// set default values for the todo
	todo.Completed = false

	todo, err := t.todoRepo.Create(todo)

	if err != nil {
		return entity.Todo{}, err
	}
	return todo, nil

}

// Delete implements TodoUsecase.
func (t *todoUsecaseImpl) Delete(id int) error {
	err := t.todoRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil

}

// GetAll implements TodoUsecase.
func (t *todoUsecaseImpl) GetAll() ([]entity.Todo, error) {
	todos, err := t.todoRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// GetByID implements TodoUsecase.
func (t *todoUsecaseImpl) GetByID(id int) (entity.Todo, error) {
	todo, err := t.todoRepo.GetByID(id)
	if err != nil {
		return entity.Todo{}, err
	}
	return todo, nil
}

// Update implements TodoUsecase.
func (t *todoUsecaseImpl) Update(todo entity.Todo) (entity.Todo, error) {
	if todo.Title == "" {
		return entity.Todo{}, apperror.ValidationError("Title ")
	}
	if todo.CreatedAt == "" {
		return entity.Todo{}, apperror.ValidationError("CreatedAt ")
	}

	todo, err := t.todoRepo.Update(todo)
	if err != nil {
		return entity.Todo{}, err
	}
	return todo, nil
}
