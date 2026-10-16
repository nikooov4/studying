package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func F() *http.Client {

	client := http.Client{}
	transport := http.Transport{}

	client.Timeout = 5 * time.Second
	transport.MaxIdleConns = 20
	transport.MaxIdleConnsPerHost = 5
	transport.IdleConnTimeout = 30 * time.Second
	transport.DisableKeepAlives = false

	client.Transport = &transport

	return &client
}

func main() {

	client := F()
	req, err := http.NewRequest(
		http.MethodGet,
		"http://localhost:8080/health",
		http.NoBody,
	)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < 5; i++ {
		func() {
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()
			data, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err, req.URL.String())
				return
			}
			fmt.Println(i+1, req.URL.String(), resp.StatusCode, len(data))
		}()
	}
}
