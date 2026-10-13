package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

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
	client.Timeout = 5 * time.Second
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			resp, err := client.Do(request)
			defer wg.Done()
			if err != nil {
				log.Println(err, request.URL.String())
				return
			}
			defer resp.Body.Close()
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err, request.URL.String())
				return
			}
			fmt.Println(request.URL.String(), resp.StatusCode, len(data))
			if resp.StatusCode > 299 || resp.StatusCode < 200 {
				fmt.Println("Status NOT Ok")
			}
		}()
	}
	wg.Wait()
}
