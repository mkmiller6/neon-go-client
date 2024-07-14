// Package client provides a Neon client for invoking APIs across all resources
package client

import (
	"github.com/mkmiller6/neon-go-client"
	"github.com/mkmiller6/neon-go-client/accounts"
	"github.com/mkmiller6/neon-go-client/eventregistrations"
	"github.com/mkmiller6/neon-go-client/events"
	"github.com/mkmiller6/neon-go-client/memberships"
)

type API struct {
	// Client for invoking /accounts APIs
	Accounts *accounts.Client
	// Client for invoking /events APIs
	Events *events.Client
	// Client for invoking /eventRegistrations APIs
	EventRegistrations *eventregistrations.Client
	// Client for invoking /memberships APIs
	Memberships *memberships.Client
}

func (a *API) init(user, apiKey string, backend neon.Backend) {
	a.Accounts = accounts.NewClient(user, apiKey, backend)
	a.Events = events.NewClient(user, apiKey, backend)
	a.EventRegistrations = eventregistrations.NewClient(user, apiKey, backend)
	a.Memberships = memberships.NewClient(user, apiKey, backend)
}

func New(user, apiKey string, backend neon.Backend) *API {
	api := API{}
	api.init(user, apiKey, backend)
	return &api
}
