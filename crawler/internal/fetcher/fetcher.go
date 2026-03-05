package fetcher

import "net/http"

// Fetcher defines the interface for fetching web pages.
type Fetcher interface {
	Fetch(url string) (*http.Response, error)
}

// DefaultFetcher is a simple implementation of Fetcher using net/http.
type DefaultFetcher struct {
	Client *http.Response
}

func (f *DefaultFetcher) Fetch(url string) (*http.Response, error) {
	return http.Get(url)
}
