// Package memberships provides the /memberships APIs
package memberships

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
