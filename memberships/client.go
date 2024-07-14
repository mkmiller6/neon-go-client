// Package memberships provides the /memberships APIs
package memberships

type Client struct {
	user, key string
}

func NewClient(user, apiKey string) *Client {
	return &Client{
		user: user,
		key:  apiKey,
	}
}
