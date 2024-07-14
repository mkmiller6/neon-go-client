// Package events provides the /events APIs
package events

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mkmiller6/neon-go-client"
)

type Client struct {
	b         neon.Backend
	user, key string
}

func NewClient(user, apiKey string, backend neon.Backend) *Client {
	return &Client{
		user: user,
		key:  apiKey,
		b:    backend,
	}
}

func (c *Client) GetByID(id int) (*neon.Event, error) {
	path := fmt.Sprintf("/v2/events/%d", id)
	event := &neon.Event{}
	resp, err := c.b.Call(http.MethodGet, path, c.user, c.key)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, event)

	if err != nil {
		return nil, err
	}

	return event, nil
}
