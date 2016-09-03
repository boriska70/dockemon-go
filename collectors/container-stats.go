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
	contData [] containerData
	dataType string
}
type containerData struct {
	id string
	names string
	image string
	labels string
	ports string
	status string
	host hostData
}
type hostData struct {
	hostName string
	os string
	memoryTotalGB int
	systemTyme string
	dockerVersion string
	totalContainers int
	runningContainers int
}

func SetConstantHostData(client doClient)  {
	info, _ := client.dc.Info(context.Background())
	Host.hostName = info.Name
	Host.os = info.OperatingSystem
	Host.memoryTotalGB = int (info.MemTotal / 1024 / 1024)
	Host.dockerVersion = info.ServerVersion
}

func (cbd *ContainersBulkData) addContainerData(cd containerData) [] containerData {
	cbd.contData = append(cbd.contData, cd)
	return cbd.contData
}

func ContainerStats(client doClient, ch chan ContainersBulkData) {
	log.Println("Collectiong containers data...")
	log.Println(client.dc.Info(context.Background()))
	info, _ := client.dc.Info(context.Background())

	var contBulk ContainersBulkData

	Host.systemTyme = info.SystemTime
	Host.totalContainers = info.Containers
	Host.runningContainers = info.ContainersRunning
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
			cont.id = c.ID
			cont.image = c.Image
//			cont.names = strings.Join(c.Names,",")
//			cont.labels = strings.Join(c.Labels,",")
//			cont.ports = strings.Join(c.Ports,",")
			cont.host = Host
			contBulk.contData = contBulk.addContainerData(cont)
		}
		if len(contBulk.contData) > 0 {
			ch <- contBulk
		}
		time.Sleep(time.Duration(client.contListIntervalSec) * time.Second)
	}
}
