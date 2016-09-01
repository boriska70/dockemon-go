package collectors

import "fmt"

import (
	"github.com/docker/engine-api/types"
	"github.com/docker/distribution/context"
	log "github.com/Sirupsen/logrus"
	"time"
)

func ContainerStats(client doClient, esClient esClient) {

	log.Println(client.dc.Info(context.Background()))
	info, _ := client.dc.Info(context.Background())
	fmt.Printf("Hostname: %v \n", info.Name)
	fmt.Printf("ContainersRunning: %v out of %v \n", info.ContainersRunning, info.Containers)
	fmt.Printf("OperatingSystem: %v \n", info.OperatingSystem)
	fmt.Printf("Memory Total in GB: %d \n", info.MemTotal / 1024 / 1024)
	fmt.Printf("SystemTime: %v \n", info.SystemTime)
	fmt.Printf("Docker version: %v \n", info.ServerVersion)

	options := types.ContainerListOptions{All:false}
	for {
		fmt.Printf("CPU Usage: %v \n", types.CPUStats{}.CPUUsage)
		fmt.Printf("Memory Usage: %v \n", types.MemoryStats{}.MaxUsage)
		containers, err := client.dc.ContainerList(context.Background(), options)
		if err != nil {
			panic(err)
		}

		for _, c := range containers {
			fmt.Println(c.ID, c.Names, c.Image, c.Labels, c.Ports, c.Status)
			//fmt.Println(c.ID, c.NetworkSettings.Networks, c.Names, c.Command, c.Created, c.Image, c.Labels, c.Ports, c.SizeRootFs, c.SizeRw, c.Status)
			//fmt.Printf("HostConfig-NetworkNode is %v \n", c.HostConfig.NetworkMode)
			//fmt.Printf("Networks: %v \n", c.NetworkSettings.Networks)


			time.Sleep(time.Duration(client.contListIntervalSec) * time.Second)
		}
	}
}
