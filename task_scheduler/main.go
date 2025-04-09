package main

import (
	"log"
	"net/http"
	"task_scheduler/api"
	"task_scheduler/queue"
)

func main() {
	taskMgr := queue.NewTaskManager()
	taskMgr.StartScheduler()
	services := &api.TaskHandler{TaskMgr: taskMgr}

	http.HandleFunc("/push", services.PushTask)
	http.HandleFunc("/consume", services.ConsumeTask)

	log.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}
