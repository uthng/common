package k8s

import (
    //"fmt"
    "github.com/pkg/errors"

    "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    //"k8s.io/client-go/kubernetes"
)

// GetPodList returns a list of pods for the given namespace without any option
// if namespace is empty, all pods in all namespaces will be returned
func (c *Client) GetPodList(namespace string) (*v1.PodList, error) {
    return c.CoreV1().Pods(namespace).List(metav1.ListOptions{})
}

// GetPodListWithOpts returns a list of pods for all namespaces
// corresponding to certains given options
func (c *Client) GetPodListWithOpts(namespace string, kind, label, field string, uninitialized, watch bool, resource string) (*v1.PodList, error) {
    opts := metav1.ListOptions{}

    if kind != "" {
        opts.TypeMeta.Kind = kind
    }

    if label != "" {
        opts.LabelSelector = label
    }

    if field != "" {
        opts.FieldSelector = field
    }

    if uninitialized == true {
        opts.IncludeUninitialized = uninitialized
    }

    if watch == true {
        opts.Watch = watch
    }

    if resource != "" {
        opts.ResourceVersion = resource
    }

    return c.CoreV1().Pods(namespace).List(opts)
}

// GetPodFromListByName returns a pod in the PodList corresponding to the given name
func (c *Client) GetPodFromListByName (name string, pods *v1.PodList) (*v1.Pod, error) {
    for _, pod := range pods.Items {
        if pod.ObjectMeta.Name == name {
            return &pod, nil
        }
    }

    return nil, errors.Errorf("No pod found with the name %s", name)
}


