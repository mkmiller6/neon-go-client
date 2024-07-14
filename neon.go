// Package neon provides the bindings for Neon REST APIs
package neon

import (
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/http"
)

type Backend interface {
	Call(method, path, user, key string) (*http.Response, error)
}

type backendImplementation struct {
	baseURL       string
	httpClient    *http.Client
	leveledLogger *slog.Logger
}

func (b *backendImplementation) Call(method, path, user, key string) (*http.Response, error) {
	authSignature := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, key)))
	authorization := "Basic " + authSignature

	url := b.baseURL + path

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorization)

	resp, err := b.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetBackendWithConfig() Backend {
	return &backendImplementation{
		baseURL:       "https://api.neoncrm.com",
		httpClient:    &http.Client{},
		leveledLogger: &slog.Logger{},
	}
}
