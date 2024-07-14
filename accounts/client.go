// Package accounts provides the /accounts APIs
package accounts

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

func (c *Client) GetByID(id int) (*neon.Account, error) {
	path := fmt.Sprintf("/v2/accounts/%d", id)
	account := &neon.Account{}
	resp, err := c.b.Call(http.MethodGet, path, c.user, c.key)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, account)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (c *Client) Search(params *neon.SearchParams) *SearchIter {
	return &SearchIter{
		SearchIter: neon.GetSearchIter(params, func(p *neon.Params) neon.SearchContainer {
			list := &neon.AccountSearchResult{}
			return list
		}),
	}
}

// Search iterator for accounts
type SearchIter struct {
	*neon.SearchIter[neon.Account]
}

// Account returns the account which the iterator is currently pointing to
func (i *SearchIter) Account() *neon.Account {
	return i.Current()
}

// AccountSearchResult returns the current list object which the iterator is currently using.
// List objects will change as new API calls are made to continue pagination.
func (i *SearchIter) AccountSearchResult() *neon.AccountSearchResult {
	return i.SearchResult().(*neon.AccountSearchResult)
}
