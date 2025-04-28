package handler

import (
	"belajar_go/apperror"
	"belajar_go/entity"
	"belajar_go/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{todoUsecase: todoUsecase}
}

func (h *TodoHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/todos", h.CreateTodo).Methods("POST")
	r.HandleFunc("/todos", h.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", h.DeleteTodo).Methods("DELETE")
	r.HandleFunc("/todos/{id}", h.GetTodoByID).Methods("GET")
	r.HandleFunc("/todos/{id}", h.UpdateTodo).Methods("PUT")
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo entity.Todo
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&todo)
	if err != nil {

		WriteJSONResponse(w, http.StatusBadRequest, "Invalid JSON Format or Field", nil)
		return
	}

	todo, err = h.todoUsecase.Create(todo)
	if err != nil {

		statusCode, message := apperror.DetermineErrorType(err)
		WriteJSONResponse(w, statusCode, message, nil)
		return
	}

	WriteJSONResponse(w, http.StatusCreated, "Todo created successfully", todo)
}

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.todoUsecase.GetAll()
	if err != nil {
		statusCode, message := apperror.DetermineErrorType(err)
		WriteJSONResponse(w, statusCode, message, nil)

		return
	}
	WriteJSONResponse(w, http.StatusOK, "Todos retrieved successfully", todos)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, "Invalid ID", nil)
		return
	}

	err = h.todoUsecase.Delete(id)
	if err != nil {
		statusCode, message := apperror.DetermineErrorType(err)
		WriteJSONResponse(w, statusCode, message, nil)
		return
	}

	WriteJSONResponse(w, http.StatusOK, "Todo deleted successfully", nil)

}

func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, "Invalid ID", nil)
		return
	}

	todo, err := h.todoUsecase.GetByID(id)
	if err != nil {
		statusCode, message := apperror.DetermineErrorType(err)
		WriteJSONResponse(w, statusCode, message, nil)
		return
	}

	WriteJSONResponse(w, http.StatusOK, "Todo retrieved successfully", todo)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, "Invalid ID", nil)
		return
	}

	var todo entity.Todo
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&todo)
	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, "Invalid JSON Format or Field", nil)
		return
	}

	todo.ID = id // set the ID to the todo struct

	todo, err = h.todoUsecase.Update(todo)
	if err != nil {
		statusCode, message := apperror.DetermineErrorType(err)
		WriteJSONResponse(w, statusCode, message, nil)
		return
	}

	WriteJSONResponse(w, http.StatusOK, "Todo updated successfully", todo)
}
