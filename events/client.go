// Package events provides the /events APIs
package events

type Client struct {
	user, key string
}

func NewClient(user, apiKey string) *Client {
	return &Client{
		user: user,
		key:  apiKey,
	}
}
