// internal/repository/user_repository.go

package repository

import (
	"io"
	"log"
	"net/http"
	"time"
)

func FetchUsers(url string) ([]byte, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			log.Println("Error closing response body:", cerr)
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
