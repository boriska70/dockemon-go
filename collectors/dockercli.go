package collectors

import (
	"github.com/docker/docker/client"
)

type doClient struct {
	dc                  *client.Client
	contListIntervalSec int64
}

// Create new docker API client based on the environment variables
func NewDockerClient(dci int64) doClient {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	} else {
		return doClient{dc: cli, contListIntervalSec: dci}
	}
}
