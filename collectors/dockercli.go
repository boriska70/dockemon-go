package collectors

import (
	"github.com/docker/docker/client"
)

type doClient struct {
	dc                  *client.Client
	contListIntervalSec int64
}

func NewDockerClient(cci int64) doClient {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	} else {
		return doClient{dc:cli, contListIntervalSec:cci}
	}
}

