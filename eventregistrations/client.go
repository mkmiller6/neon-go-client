// Package eventregistrations provides the /eventRegistrations APIs
package eventregistrations

type Client struct {
	user, key string
}

func NewClient(user, apiKey string) *Client {
	return &Client{
		user: user,
		key:  apiKey,
	}
}
