package k8s

import (

    "k8s.io/client-go/kubernetes"
)

// Client is en embedded type of kubernetes clientset
type Client struct {
    *kubernetes.Clientset
}

// NewClient initializes a new clientset for the given config
func NewClient(config *Config) (*Client, error) {
    client := &Client {}

    clientset, err := kubernetes.NewForConfig(config.Config)
    if err != nil {
        return nil, err
    }

    client.Clientset = clientset

    return client, nil
}
