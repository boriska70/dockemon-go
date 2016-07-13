package collectors

import (
	"fmt"
	"github.com/samalba/dockerclient"
	log "github.com/Sirupsen/logrus"
	"time"
)

type doClient struct {
	Cl *dockerclient.DockerClient
	contListIntervalSec int64
}

type doMoClient interface {

	listContainers()

}

func NewClient(cci int64) doClient {
	cli, err := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	if err != nil {
		panic(err)
	} else {
		return doClient{Cl:cli, contListIntervalSec:cci}
	}
}

func (client doClient) ListContainers() {
	for {
		containers, err := client.Cl.ListContainers(true, false, "")
		if err != nil {
			panic(err)
		}
		log.Println(client.Cl.Info())
		for _, c := range containers {
			fmt.Println(c.Id, c.NetworkSettings.Networks, c.Names, c.Command, c.Created, c.Image, c.Labels, c.Ports, c.SizeRootFs, c.SizeRw, c.Status)
			info, _ := client.Cl.InspectContainer(c.Id)
			log.Println(info.HostConfig.Ulimits)
			log.Println(info.NetworkSettings.IPAddress)
		}
		time.Sleep(time.Duration(client.contListIntervalSec) * time.Second)
	}
}

func (client doClient) EventCallBack(event *dockerclient.Event, ec chan error, args ...interface{}) {
	log.Printf("Received event: %#v\n", *event)

}

