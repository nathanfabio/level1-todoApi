package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nathanfabio/level1-todoApi/internal/handler"
	"github.com/nathanfabio/level1-todoApi/internal/repository"
	"github.com/nathanfabio/level1-todoApi/pkg/db"
)


func main() {
	database := db.Connect()
	repo := repository.NewTaskRepository(database)
	taskHandler := handler.NewTaskHandler(repo)

	r := chi.NewRouter()

	//Test
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Post("/register", taskHandler.RegisterHandler)
	r.Post("/login", taskHandler.LoginHandler)
	r.Post("/tasks", taskHandler.CreateTask)
	r.Get("/tasks", taskHandler.ListTasks)
	r.Put("/tasks/{id}/status", taskHandler.UpdateTaskStatus)
	r.Delete("/tasks/{id}", taskHandler.DeleteTask)


	log.Println("Starting server on :7070")
	http.ListenAndServe(":7070", r)
}