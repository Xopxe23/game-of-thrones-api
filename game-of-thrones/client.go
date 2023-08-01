package gameofthrones

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	Client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}
	return &Client{
		Client: &http.Client{
			Timeout: timeout,
			Transport: loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c *Client) GetHouses() ([]House, error) {
	resp, err := c.Client.Get("https://api.gameofthronesquotes.xyz/v1/houses")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var houses []House
	if err := json.Unmarshal(body, &houses); err != nil {
		return nil, err
	}
	return houses, nil
}
