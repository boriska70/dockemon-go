package collectors

import "fmt"

import (
	"github.com/docker/engine-api/types"
	"github.com/docker/distribution/context"
	log "github.com/Sirupsen/logrus"
	"time"
	"strings"
)

var Host hostData

type ContainersBulkData struct {
	ContData [] containerData
	DataType string
	CollectionTime time.Time
}
type containerData struct {
	Id     string
	Names  string
	Image  string
	Labels string
	Ports  [] types.Port
	Status string
	Host   hostData
}
type hostData struct {
	HostName          string
	Os                string
	MemoryTotalGB     int
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
	contBulk.DataType="container_monitor"

	options := types.ContainerListOptions{All:false}
	for {
		fmt.Printf("CPU Usage: %v \n", types.CPUStats{}.CPUUsage)
		fmt.Printf("Memory Usage: %v \n", types.MemoryStats{}.MaxUsage)
		contBulk.CollectionTime=time.Now()
		Host.TotalContainers = info.Containers
		Host.RunningContainers = info.ContainersRunning

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
			cont.Status = c.Status
			cont.Names = strings.Join(c.Names,",")
			cont.Labels = strings.Join(mapToArray(c.Labels),",")
			cont.Ports = c.Ports
			cont.Host = Host
			contBulk.ContData = contBulk.addContainerData(cont)
		}
		if len(contBulk.ContData) > 0 {
			ch <- contBulk
		}
		time.Sleep(time.Duration(client.contListIntervalSec) * time.Second)
	}
}

func mapToArray(m map[string]string) [] string  {

	res := make([]string, 0)

	for ind,val := range m {
		res = append(res,ind+"="+val)
	}

	return res
}


