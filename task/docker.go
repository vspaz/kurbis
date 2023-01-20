package task

import (
	"context"
	"github.com/docker/docker/api/types"
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
	return Result{Error: nil}
}
