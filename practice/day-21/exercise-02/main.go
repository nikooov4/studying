package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Slow(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("done"))
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {

	back := context.Background()

	ctrlC := os.Interrupt
	ctx, stop := signal.NotifyContext(back, ctrlC)
	defer stop()

	router := http.NewServeMux()
	router.HandleFunc("GET /health", Health)
	router.HandleFunc("GET /slow", Slow)
	server := http.Server{}
	server.Addr = "localhost:8081"
	server.Handler = router
	chErr := make(chan error, 1)
	go func() {
		fmt.Println("starting server")
		err := server.ListenAndServe()
		chErr <- err
	}()
	select {
	case <-ctx.Done():
		ctxFiveSec, cancelFiveSec := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFiveSec()
		fmt.Println("server ending")
		errServer := server.Shutdown(ctxFiveSec)
		if errServer != nil {
			fmt.Println(errServer)
			return
		}
		fmt.Println("server end")
	case err := <-chErr:
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println(err)
		} else {
			fmt.Printf("foreing err: %v", err)
			return
		}
	}
}
