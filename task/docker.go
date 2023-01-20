package task

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Docker struct {
	Client      *client.Client
	Config      Config
	ContainerId string
	Logger      *logrus.Logger
}

type Result struct {
	Error       error
	Action      string
	ContainerId string
	Result      string
}

func (d *Docker) Run() Result {
	ctx := context.Background()
	reader, err := d.Client.ImagePull(
		ctx,
		d.Config.Image,
		types.ImagePullOptions{},
	)
	if err != nil {
		d.Logger.Errorf("Error pulling image %s %s", d.Config.Image, err)
		return Result{Error: err}
	}
	io.Copy(os.Stdout, reader)

	restartPolicy := container.RestartPolicy{
		Name: d.Config.RestartPolicy,
	}

	resources := container.Resources{
		Memory: d.Config.Memory,
	}

	containerConfig := container.Config{
		Image: d.Config.Image,
		Env:   d.Config.Env,
	}

	hostConfig := container.HostConfig{
		RestartPolicy:   restartPolicy,
		Resources:       resources,
		PublishAllPorts: true,
	}

	resp, err := d.Client.ContainerCreate(
		ctx,
		&containerConfig,
		&hostConfig,
		nil,
		nil,
		d.Config.Name,
	)
	if err != nil {
		d.Logger.Infof("Error creating container from image %s: %v", d.Config.Image, err)
		return Result{Error: err}
	}

	err := d.Client.ContainerStart(
		ctx, resp.ID, types.ContainerStartOptions{},
	)
	if err != nil {
		d.Logger.Errorf("Error starting container %s: %v\n", resp.ID, err)
		return Result{Error: err}
	}

	return Result{Error: nil}
}
