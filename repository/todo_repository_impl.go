package repository

import (
	"belajar_go/apperror"
	"belajar_go/domain"
	"belajar_go/entity"
	"database/sql"
)

// todoRepositoryImpl gunanya untuk mengimplementasikan interface TodoRepository
type todoRepositoryImpl struct {
	db *sql.DB
}

// fungsi NewRepositoryImpl untuk membuat instance dari todoRepositoryImpl
func NewRepositoryImpl(db *sql.DB) domain.TodoRepository {
	return &todoRepositoryImpl{
		db: db,
	}
}

// Create implements domain.TodoRepository.
func (t *todoRepositoryImpl) Create(todo entity.Todo) (entity.Todo, error) {
	result, err := t.db.Exec("INSERT INTO todos(title, completed, created_at) VALUES (?, ?, ?)", todo.Title, todo.Completed, todo.CreatedAt)

	// check if the insert was successful
	if err != nil {
		return entity.Todo{}, err
	}

	id, err := result.LastInsertId() // get the last inserted id
	if err != nil {
		return entity.Todo{}, err
	}

	todo.ID = int(id) // set the id to the todo struct
	return todo, nil

}

// Delete implements domain.TodoRepository.
func (t *todoRepositoryImpl) Delete(id int) error {
	del, err := t.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}

	// check if the delete was successful
	rowsAffected, err := del.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apperror.ErrNotFound // return an error if no rows were affected
	}
	return nil // return nil if the delete was successful
}

// GetAll implements domain.TodoRepository.
func (t *todoRepositoryImpl) GetAll() ([]entity.Todo, error) {
	get, err := t.db.Query("SELECT id, title, completed, created_at FROM todos")

	if err != nil {
		return nil, err
	}

	defer get.Close() // close the connection after use

	var todos []entity.Todo
	for get.Next() {
		var todo entity.Todo
		if err := get.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
			&todo.CreatedAt,
		); err != nil {
			return nil, err
		}
		todos = append(todos, todo) // append the todo to the slice

	}
	return todos, nil // return the slice of todos

}

// GetByID implements domain.TodoRepository.
func (t *todoRepositoryImpl) GetByID(id int) (entity.Todo, error) {
	var todo entity.Todo
	err := t.db.QueryRow("SELECT id, title, completed, created_at FROM todos WHERE id = ?", id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Todo{}, apperror.ErrNotFound // return an error if no rows were found
		}
		return entity.Todo{}, err
	}
	return todo, nil // return the todo if found

}

// Update implements domain.TodoRepository.
func (t *todoRepositoryImpl) Update(todo entity.Todo) (entity.Todo, error) {
	update, err := t.db.Exec("UPDATE todos SET title = ?, completed = ?, created_at = ? WHERE id = ?", todo.Title, todo.Completed, todo.CreatedAt, todo.ID)
	if err != nil {
		return entity.Todo{}, err
	}
	// check if the update was successful
	rowsAffected, err := update.RowsAffected()
	if err != nil {
		return entity.Todo{}, err
	}
	if rowsAffected == 0 {
		return entity.Todo{}, apperror.ErrNotFound // return an error if no rows were affected
	}
	return todo, nil // return the updated todo
}
