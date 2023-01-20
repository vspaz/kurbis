package task

import "github.com/docker/docker/client"

type Docker struct {
	Client      *client.Client
	Config      Config
	ContainerId string
}

type Result struct {
	Error       error
	Action      string
	ContainerId string
	Result      string
}
