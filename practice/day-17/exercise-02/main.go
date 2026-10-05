package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	uuid "github.com/google/uuid"
)

type Task struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Importance  string    `json:"importance"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
}

type Store struct {
	tasks map[uuid.UUID]Task
	mu    sync.Mutex
}

func NewStore() *Store {
	tasks := make(map[uuid.UUID]Task)
	return &Store{
		tasks: tasks,
	}
}

func (s *Store) Add(task Task) error {
	s.mu.Lock()
	if _, ok := s.tasks[task.UUID]; ok {
		s.mu.Unlock()
		return errors.New("UUID занят")
	}
	s.tasks[task.UUID] = task
	s.mu.Unlock()

	return nil
}

func (s *Store) Get(id uuid.UUID) (Task, error) {
	s.mu.Lock()
	if _, ok := s.tasks[id]; !ok {
		s.mu.Unlock()
		return Task{}, errors.New("Not found")
	}
	defer s.mu.Unlock()
	return s.tasks[id], nil
}

func (s *Store) List() []Task {
	taskArr := make([]Task, 0)
	s.mu.Lock()
	for _, v := range s.tasks {
		taskArr = append(taskArr, v)
	}
	s.mu.Unlock()
	return taskArr
}

func (s *Store) Delete(id uuid.UUID) error {
	s.mu.Lock()
	if _, ok := s.tasks[id]; !ok {
		s.mu.Unlock()
		return errors.New("Not found for Delete")
	}
	delete(s.tasks, id)
	s.mu.Unlock()
	return nil
}

func (s *Store) Update(task Task) error {
	s.mu.Lock()
	if _, ok := s.tasks[task.UUID]; !ok {
		s.mu.Unlock()
		return errors.New("Not found")
	}
	s.tasks[task.UUID] = task
	s.mu.Unlock()
	return nil
}

func (s *Store) CreateHTTP(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "error decode", 400)
		return
	}
	id := uuid.New()
	task.UUID = id
	task.Done = false
	errAdd := s.Add(task)
	if errAdd != nil {
		http.Error(w, "error save", 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
	err1 := json.NewEncoder(w).Encode(task)
	if err1 != nil {
		log.Print(err1)
	}
}

func (s *Store) ListHTTP(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	tasks = s.List()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err1 := json.NewEncoder(w).Encode(tasks)
	if err1 != nil {
		log.Print(err1)
	}
}

func (s *Store) ListHTTPID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	uid, errID := uuid.Parse(id)
	if errID != nil {
		http.Error(w, "incorrect uuid", 400)
		return
	}
	task, errGet := s.Get(uid)
	if errGet != nil {
		http.Error(w, "Not Found", 404)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err1 := json.NewEncoder(w).Encode(task)
	if err1 != nil {
		log.Print(err1)
	}
}

func (s *Store) DeleteHTTP(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	uid, errID := uuid.Parse(id)
	if errID != nil {
		http.Error(w, "incorrect uuid", 400)
		return
	}
	errDelete := s.Delete(uid)
	if errDelete != nil {
		http.Error(w, "error delete", 404)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Store) PutHTTP(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	uid, errID := uuid.Parse(id)
	if errID != nil {
		http.Error(w, "incorrect uuid", 400)
		return
	}
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "error decode", 400)
		return
	}
	task.UUID = uid
	errUpdate := s.Update(task)
	if errUpdate != nil {
		http.Error(w, "error update", 404)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	errEncode := json.NewEncoder(w).Encode(task)
	if errEncode != nil {
		log.Printf("encode error: %v", errEncode)
		return
	}
}


func main() {

	store := NewStore()

	router := http.NewServeMux()
	router.HandleFunc("POST /tasks", store.CreateHTTP)
	router.HandleFunc("GET /tasks", store.ListHTTP)
	router.HandleFunc("GET /tasks/{id}", store.ListHTTPID)
	router.HandleFunc("DELETE /tasks/{id}", store.DeleteHTTP)
	router.HandleFunc("PUT /tasks/{id}", store.PutHTTP)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
