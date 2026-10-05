package main

import (
	"slices"
	"testing"

	"github.com/google/uuid"
)

func TestAdd(t *testing.T) {

	id := uuid.New()
	task := Task{
		UUID: id,
		Name: "GetAll",
	}
	store := NewStore()
	err := store.Add(task)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	task1, err := store.Get(id)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if task1 != task {
		t.Fatal("Error Get")
	}

	task2 := Task{
		UUID: id,
		Name: "GetNOOO",
	}
	err2 := store.Add(task2)
	if err2 == nil {
		t.Fatalf("error, shoul be error")
	}
	task_test, err_test := store.Get(id)

	if err_test != nil {
		t.Fatalf("error: %v", err_test)
	}
	if task != task_test {
		t.Fatal("Error task")
	}

	id2 := uuid.New()
	task3 := Task{
		UUID: id2,
		Name: "GetAll",
	}
	err4 := store.Add(task3)
	if err4 != nil {
		t.Fatalf("error: %v", err4)
	}
}

func TestGet(t *testing.T) {

	id := uuid.New()
	store := NewStore()
	_, err := store.Get(id)

	if err == nil {
		t.Fatalf("error: %v", err)
	}
	task := Task{
		UUID: id,
		Name: "GetAll",
	}
	err1 := store.Add(task)
	if err1 != nil {
		t.Fatalf("error: %v", err1)
	}

	task2, err2 := store.Get(id)
	if err2 != nil {
		t.Fatalf("error: %v", err2)
	}
	if task2 != task {
		t.Fatalf("error: get")
	}
	task2.Name = "123"
	task3, err2 := store.Get(id)
	if err2 != nil {
		t.Fatalf("error: %v", err2)
	}
	if task3 != task {
		t.Fatalf("error: get")
	}
}

func TestList(t *testing.T) {

	store := NewStore()
	tasks := store.List()
	if len(tasks) != 0 || tasks == nil {
		t.Fatal("error len")
	}

	id1 := uuid.New()
	task1 := Task{
		UUID: id1,
		Name: "GetAll",
	}
	id2 := uuid.New()
	task2 := Task{
		UUID: id2,
		Name: "GetAll",
	}
	err1 := store.Add(task1)
	if err1 != nil {
		t.Fatalf("error: %v", err1)
	}
	err2 := store.Add(task2)
	if err2 != nil {
		t.Fatalf("error: %v", err2)
	}

	tasks1 := store.List()
	if len(tasks1) != 2 {
		t.Fatal("error list")
	}
	exists1 := slices.Contains(tasks1, task1)
	exists2 := slices.Contains(tasks1, task2)
	if !exists1 || !exists2 {
		t.Fatal("error list")
	}
	tasks11 := tasks1[0]
	tasks1[0].Name = "123"
	tasks111, err := store.Get(tasks11.UUID)
	if err != nil {
		t.Fatal("error")
	}
	if tasks111 != tasks11 {
		t.Fatal("error storage")
	}
}

func TestDelete(t *testing.T) {

	store := NewStore()
	id1 := uuid.New()
	task1 := Task{
		UUID: id1,
		Name: "GetAll",
	}
	id2 := uuid.New()
	task2 := Task{
		UUID: id2,
		Name: "GetAll",
	}
	err1 := store.Add(task1)
	if err1 != nil {
		t.Fatalf("error: %v", err1)
	}
	err2 := store.Add(task2)
	if err2 != nil {
		t.Fatalf("error: %v", err2)
	}
	err3 := store.Delete(id2)
	if err3 != nil {
		t.Fatalf("error: delete %v", err3)
	}

	_, err4 := store.Get(id2)
	if err4 == nil {
		t.Fatalf("error: %v", err4)
	}
	task5, err5 := store.Get(id1)
	if err5 != nil {
		t.Fatalf("error: %v", err5)
	}
	if task5 != task1 {
		t.Fatal("errro get")
	}
	err6 := store.Delete(id2)
	if err6 == nil {
		t.Fatalf("error: delete %v", err6)
	}
}
