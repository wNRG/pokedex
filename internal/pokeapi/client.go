package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client {
		httpClient: &http.Client{},
	}
}

func (c *Client) get(url string, destination any) error {
	res, err := c.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("Request failed with status code: %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(destination); err != nil {
		return err
	}

	return nil
}
