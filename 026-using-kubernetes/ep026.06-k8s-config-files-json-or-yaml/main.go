package main

import (
	"encoding/json"
	"fmt"
)

type Resource struct {
	ApiVersion string   `json:"apiVersion,omitempty"`
	Kind       string   `json:"kind,omitempty"`
	Metadata   Metadata `json:"metadata,omitempty"`
	Spec       Spec     `json:"spec,omitempty"`
}

type Metadata struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
}

type Spec struct {
	Containers []Container `json:"containers"`
}

type Container struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

func main() {
	pod := Resource{
		ApiVersion: "v1",
		Kind:       "Pod",
		Metadata: Metadata{
			Name: "my-stack",
			Labels: map[string]string{
				"frontend": "production",
				"location": "east",
			},
		},
		Spec: Spec{
			Containers: []Container{
				{
					Name:  "redis",
					Image: "redis",
				},
				{
					Name:  "server",
					Image: "server-02",
				},
			},
		},
	}

	buf, _ := json.Marshal(pod)
	fmt.Printf("%s\n", buf)
}
