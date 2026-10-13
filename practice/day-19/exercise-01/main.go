package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func handl(w http.ResponseWriter, r *http.Request) {

}

func main() {


	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "GET /health", &io.LimitedReader{})
	client := http.Client{}
	client.Timeout = time.Second * 5
	

}
