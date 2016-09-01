package collectors

import (
    "fmt"
    "github.com/docker/engine-api/client"
    "github.com/docker/engine-api/types"
    "github.com/docker/distribution/context"
    log "github.com/Sirupsen/logrus"
    "time"
)

type doClient struct {
    dc *client.Client
    contListIntervalSec int64
}

type doMoClient interface {
    listContainers()
}

func NewClient(cci int64) doClient {
    //defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
    //cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, defaultHeaders)
    cli, err := client.NewEnvClient()
    if err != nil {
	panic(err)
    } else {
	return doClient{dc:cli, contListIntervalSec:cci}
    }
}

func (client doClient) ListContainers() {
    options := types.ContainerListOptions{All:true}

    for {

	containers, err := client.dc.CheckpointList(context.Background(), options)
	if err != nil {
	    panic(err)
	}
	log.Println(client.dc.Info())
	for _, c := range containers {
	    fmt.Println(c.ID)

/*

	    fmt.Println(c..Id, c.NetworkSettings.Networks, c.Names, c.Command, c.Created, c.Image, c.Labels, c.Ports, c.SizeRootFs, c.SizeRw, c.Status)
	    info, _ := client.dc.InspectContainer(c.Id)
	    log.Println(info.HostConfig.Ulimits)
	    log.Println(info.NetworkSettings.IPAddress)*/
	}
	time.Sleep(time.Duration(client.contListIntervalSec) * time.Second)
    }
}

func (client doClient) EventCallBack(event *dockerclient.Event, ec chan error, args ...interface{}) {
    log.Printf("Received event: %#v\n", *event)

}

