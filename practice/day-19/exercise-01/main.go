package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func handl(w http.ResponseWriter, r *http.Request) {

}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"http://localhost:8080/health",
		http.NoBody)

	if err != nil {
		log.Println(err)
		return
	}
	client := http.Client{}
	client.Timeout = time.Second * 5
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		fmt.Println("Status NOT Ok")
	}
	fmt.Println(resp.StatusCode)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(len(data))
}
