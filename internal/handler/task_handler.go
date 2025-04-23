package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nathanfabio/level1-todoApi/internal/model"
	"github.com/nathanfabio/level1-todoApi/internal/repository"
)

type statusUpdate struct {
	Done bool `json:"done"`
}

type TaskHandler struct {
	Repo *repository.TaskRepository
}

func NewTaskHandler(repo *repository.TaskRepository) *TaskHandler {
	return &TaskHandler{Repo: repo}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Create(&task); err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		log.Printf("Error creating task: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) ListTasks(w http.ResponseWriter, r *http.Request) {
	doneParam := r.URL.Query().Get("done")

	if doneParam != "" {
		done, err := strconv.ParseBool(doneParam)
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}

		tasks, err := h.Repo.FindByStatus(done)
		if err != nil {
			http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(tasks)
		return
	}

	
	tasks, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)

}

func (h *TaskHandler) UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var input statusUpdate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.Repo.UpdateStatus(id, input.Done); err != nil {
		http.Error(w, "Failed to update task status", http.StatusInternalServerError)
		log.Printf("Error updating task status: %v", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})

}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		log.Printf("Error deleting task: %v", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

