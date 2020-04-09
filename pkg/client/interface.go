package client

import "k8s.io/client-go/rest"

// APIClient is the interface implemented by the Alauda API client.
type APIClient interface {
	Config() *rest.Config
}

// Type checking to ensure Client correctly implements AlaudaClient.
var _ APIClient = &Client{}
