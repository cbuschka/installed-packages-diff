package _package

import (
  "fmt"
  "github.com/cbuschka/pkgdiff/internal/server"
)

var providers = []Provider{
	Provider(&Dnf{}),
}

func DetectProvider(channel server.Channel) (Provider, error) {

	for _, provider := range providers {
		found, err := provider.IsAvailable(channel)
		if err != nil {
			return nil, err
		}

		if found {
			return provider, nil
		}
	}

	return nil, fmt.Errorf("no known package provider detected")
}
