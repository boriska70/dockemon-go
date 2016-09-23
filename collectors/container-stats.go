package collectors


import (
	"github.com/docker/engine-api/types"
	"github.com/docker/distribution/context"
	log "github.com/Sirupsen/logrus"
	"time"
	"strings"
)

var HostForContainers hostData

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
	Host *HostStaticData
	TotalContainers int
	RunningContainers int
}

func (cbd *ContainersBulkData) addContainerData(cd containerData) [] containerData {
	cbd.ContData = append(cbd.ContData, cd)
	return cbd.ContData
}

func ContainerStats(client doClient, ch chan ContainersBulkData) {
	log.Println("Collecting containers data...")
	log.Println(client.dc.Info(context.Background()))
	info, _ := client.dc.Info(context.Background())

	HostForContainers.Host = getHostStaticData()
	HostForContainers.TotalContainers = info.Containers
	HostForContainers.RunningContainers = info.ContainersRunning
	log.Info("Found ",HostForContainers.TotalContainers," containers, of them running: ", HostForContainers.RunningContainers)

	var contBulk ContainersBulkData
	contBulk.DataType="container_monitor"

	options := types.ContainerListOptions{All:false}
	for {
		contBulk.CollectionTime=time.Now()
		containers, err := client.dc.ContainerList(context.Background(), options)
		if err != nil {
			panic(err)
		}

		for _, c := range containers {
			log.Debug("Container found: ",c.ID, c.Names, c.Image, c.Labels, c.Ports, c.Status)
			//fmt.Println(c.ID, c.NetworkSettings.Networks, c.Names, c.Command, c.Created, c.Image, c.Labels, c.Ports, c.SizeRootFs, c.SizeRw, c.Status)
			//fmt.Printf("HostConfig-NetworkNode is %v \n", c.HostConfig.NetworkMode)
			//fmt.Printf("Networks: %v \n", c.NetworkSettings.Networks)
			var cont containerData
			cont.Id = c.ID
			cont.Image = c.Image
			cont.Status = c.Status
			cont.Names = strings.Join(c.Names,",")
			cont.Labels = strings.Join(MapToArray(c.Labels),",")
			cont.Ports = c.Ports
			cont.Host = HostForContainers
			contBulk.ContData = contBulk.addContainerData(cont)
		}
		if len(contBulk.ContData) > 0 {
			ch <- contBulk
		}
		log.Debug("Container monitor is going to sleep under the next collection")
		time.Sleep(time.Duration(client.contListIntervalSec) * time.Second)
	}
}
