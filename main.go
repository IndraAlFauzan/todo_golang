package main

import (
	"belajar_go/handler"
	"belajar_go/middleware"
	"belajar_go/repository"
	"belajar_go/usecase"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/todo_db")
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Database unreachable:", err)
	}

	todoRepo := repository.NewRepositoryImpl(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepo)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	r := mux.NewRouter()
	todoHandler.RegisterRoutes(r)

	r.MethodNotAllowedHandler = middleware.MethodNotAllowedHandler()

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
