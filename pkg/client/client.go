package client

import (
	"flag"
	"os"
	"path/filepath"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Client is the API client for the Alauda platform.
type Client struct {
	config *rest.Config
}

// NewClient creates a new Alauda API client.
func NewClient() *Client {
	return &Client{}
}

// Config field.
func (client *Client) Config() *rest.Config {
	return client.config
}

// Initialize should be called before using the client.
func (client *Client) Initialize() error {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return err
	}

	client.config = config

	return nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
