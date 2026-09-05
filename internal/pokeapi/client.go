package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/wNRG/pokedex/internal/pokeapi/pokecache"
	"time"
)

type Client struct {
	httpClient *http.Client
	cache *pokecache.Cache
}

func NewClient(timeout, interval time.Duration) *Client {
	return &Client {
		httpClient: &http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
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
