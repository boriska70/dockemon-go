package collectors

import "fmt"

import (
	"github.com/docker/engine-api/types"
	"github.com/docker/distribution/context"
	log "github.com/Sirupsen/logrus"
	"time"
_	"strings"
)

var Host hostData

type ContainersBulkData struct {
	ContData [] containerData
	DataType string
}
type containerData struct {
	Id     string
	Names  string
	Image  string
	Labels string
	Ports  string
	Status string
	Host   hostData
}
type hostData struct {
	HostName          string
	Os                string
	MemoryTotalGB     int
	SystemTyme        string
	DockerVersion     string
	TotalContainers   int
	RunningContainers int
}

func SetConstantHostData(client doClient)  {
	info, _ := client.dc.Info(context.Background())
	Host.HostName = info.Name
	Host.Os = info.OperatingSystem
	Host.MemoryTotalGB = int (info.MemTotal / 1024 / 1024)
	Host.DockerVersion = info.ServerVersion
}

func (cbd *ContainersBulkData) addContainerData(cd containerData) [] containerData {
	cbd.ContData = append(cbd.ContData, cd)
	return cbd.ContData
}

func ContainerStats(client doClient, ch chan ContainersBulkData) {
	log.Println("Collectiong containers data...")
	log.Println(client.dc.Info(context.Background()))
	info, _ := client.dc.Info(context.Background())

	var contBulk ContainersBulkData

	Host.SystemTyme = info.SystemTime
	Host.TotalContainers = info.Containers
	Host.RunningContainers = info.ContainersRunning
/*	fmt.Printf("Hostname: %v \n", info.Name)
	fmt.Printf("ContainersRunning: %v out of %v \n", info.ContainersRunning, info.Containers)
	fmt.Printf("OperatingSystem: %v \n", info.OperatingSystem)
	fmt.Printf("Memory Total in GB: %d \n", info.MemTotal / 1024 / 1024)
	fmt.Printf("SystemTime: %v \n", info.SystemTime)
	fmt.Printf("Docker version: %v \n", info.ServerVersion)*/

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
			var cont containerData
			cont.Id = c.ID
			cont.Image = c.Image
//			cont.names = strings.Join(c.Names,",")
//			cont.labels = strings.Join(c.Labels,",")
//			cont.ports = strings.Join(c.Ports,",")
			cont.Host = Host
			contBulk.ContData = contBulk.addContainerData(cont)
		}
		if len(contBulk.ContData) > 0 {
			ch <- contBulk
		}
		time.Sleep(time.Duration(client.contListIntervalSec) * time.Second)
	}
}
