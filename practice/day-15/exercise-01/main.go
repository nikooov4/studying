package main

import (
	"context"
)

func waitForResult(
	ctx context.Context,
	result <-chan string,
) (string, error) {

	select {
	case v := <-result:
		return v, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func main() {

}
