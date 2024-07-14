// Package eventregistrations provides the /eventRegistrations APIs
package eventregistrations

import "github.com/mkmiller6/neon-go-client"

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
