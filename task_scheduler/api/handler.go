package api

import (
	"encoding/json"
	"net/http"
	"task_scheduler/queue"
	"time"

	"github.com/google/uuid"
)

type TaskHandler struct {
	TaskMgr *queue.TaskManager
}

type PushRequest struct {
	Payload     string  `json:"payload"`
	Delay       int     `json:"delay"` // seconds
}

type NoTaskFoundResponse struct {
	Message string `json:"message"`
	
}

func (this_handler *TaskHandler) PushTask(w http.ResponseWriter, r *http.Request) {
	var req PushRequest
	if err := json.NewDecoder((r.Body)).Decode(&req); err != nil {
		http.Error(w, "invalue request", http.StatusBadRequest)
	}

	task:= queue.Task{
		ID: uuid.NewString(),
		Payload: req.Payload,
		CreatedAt: time.Now(),
		AvailableAt: time.Now().Add(time.Duration(req.Delay)* time.Second),
	}

	this_handler.TaskMgr.AddTask(task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}


func (this_handler *TaskHandler) ConsumeTask(w http.ResponseWriter, r *http.Request) {
	task, ok := this_handler.TaskMgr.GetReadyTask()

	if !ok {
		response := NoTaskFoundResponse{
			Message: "No task ready",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}