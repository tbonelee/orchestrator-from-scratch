package worker

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
	"orchestrator-from-scratch/stats"
	"orchestrator-from-scratch/task"
)

func (a *Api) StartTaskHandler(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	te := task.TaskEvent{}
	err := d.Decode(&te)
	if err != nil {
		msg := fmt.Sprintf("Error unmarshalling body: %v\n", err)
		log.Printf(msg)
		w.WriteHeader(http.StatusBadRequest)
		e := ErrResponse{
			HTTPStatusCode: http.StatusBadRequest,
			Message:        msg,
		}
		json.NewEncoder(w).Encode(e)
		return
	}

	a.Worker.AddTask(te.Task)
	log.Printf("Added task %v\n", te.Task.ID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(te.Task)
}

func (a *Api) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(a.Worker.GetTasks())
}

func (a *Api) StopTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "taskID")
	if taskID == "" {
		log.Printf("No taskID passed in request.\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tID, _ := uuid.Parse(taskID)
	_, err := a.Worker.Db.Get(tID.String())
	if err != nil {
		log.Printf("No task with ID %v found", tID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	result, err := a.Worker.Db.Get(tID.String())
	if err != nil {
		log.Printf("Error getting task: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	taskToStop, ok := result.(*task.Task)
	if !ok {
		log.Printf("Error getting task: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	taskCopy := *taskToStop
	taskCopy.State = task.Completed
	a.Worker.AddTask(taskCopy)

	log.Printf("Added task %v to stop container %v\n", taskToStop.ID, taskToStop.ContainerID)
	w.WriteHeader(http.StatusNoContent)

}

func (a *Api) GetStatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if a.Worker.Stats != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(a.Worker.Stats)
		return
	}

	w.WriteHeader(http.StatusOK)
	stats := stats.GetStats()
	json.NewEncoder(w).Encode(stats)
}
