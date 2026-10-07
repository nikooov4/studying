package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"
)

type Js struct {
	Name       string `json:"name"`
	Importance int    `json:"importance"`
}

func Logging(next http.Handler) http.Handler {
	wrapped := func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		t := time.Since(now)
		log.Printf("method: %s, URL: %s, time: %v", r.Method, r.URL.Path, t)
	}
	return http.HandlerFunc(wrapped)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func Recovery(next http.Handler) http.Handler {
	recov := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			result := recover()
			if result != nil {
				http.Error(w, "error", 500)
				log.Print(result)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(recov)
}

func Panic(w http.ResponseWriter, r *http.Request) {
	panic("paniccc")
}

func Validate(js Js) error {
	jsTrim := strings.TrimSpace(js.Name)
	if len(jsTrim) <= 0 {
		return errors.New("invalid name")
	}
	if js.Importance < 1 || js.Importance > 5 {
		return errors.New("invalid importance")
	}
	return nil
}

func ValidateHHTP(w http.ResponseWriter, r *http.Request) {
	var js Js
	errDecode := json.NewDecoder(r.Body).Decode(&js)
	if errDecode != nil {
		// log.Fatalf("error decode: %v", errDecode)
		w.WriteHeader(400)
		return
	}
	err := Validate(js)
	if err != nil {
		w.WriteHeader(400)
		// log.Fatalf("error validate: %v", err)
		return
	}
	w.WriteHeader(204)
}

func main() {

	router := http.NewServeMux()

	health := http.HandlerFunc(Health)
	p := http.HandlerFunc(Panic)

	logPan := Logging(p)
	recPan := Recovery(logPan)

	logHeal := Logging(health)
	recHeal := Recovery(logHeal)

	val := http.HandlerFunc(ValidateHHTP)
	logVal := Logging(val)
	recVal := Recovery(logVal)

	router.Handle("GET /health", recHeal)
	router.Handle("GET /panic", recPan)
	router.Handle("POST /validate", recVal)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
