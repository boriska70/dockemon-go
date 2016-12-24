package collectors

import (
	"context"
	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"strings"
	"time"
)

// HostForContainers variable keeps the host level data for the host on which containers are been monitored.
var HostForContainers hostData

// ContainersBulkData type keeps container related data for all containers in the given host.
type ContainersBulkData struct {
	ContData          []containerData
	DataType          string
	CollectionTime    time.Time
	Host              hostData
	TotalContainers   int
	RunningContainers int
}
type containerData struct {
	Id     string
	Names  string
	Image  string
	Labels string
	Ports  []types.Port
	Status string
}
type hostData struct {
	Host *HostStaticData
}

func (cbd *ContainersBulkData) addContainerData(cd containerData) []containerData {
	cbd.ContData = append(cbd.ContData, cd)
	return cbd.ContData
}

// ContainerStats is the endless loop to collect container related data
func ContainerStats(client doClient, ch chan ContainersBulkData) {

	log.Println("Collecting containers data...")
	var contBulk ContainersBulkData
	contBulk.DataType = "container_monitor"

	HostForContainers.Host = getHostStaticData()

	for {
		info, _ := client.dc.Info(context.Background())
		options := types.ContainerListOptions{All: false}

		contBulk.CollectionTime = time.Now()
		contBulk.TotalContainers = info.Containers
		contBulk.RunningContainers = info.ContainersRunning
		log.Info("Found ", contBulk.TotalContainers, " containers, of them running: ", contBulk.RunningContainers)

		containers, err := client.dc.ContainerList(context.Background(), options)
		if err != nil {
			panic(err)
		}

		for _, c := range containers {
			log.Debug("Container found: ", c.ID, c.Names, c.Image, c.Labels, c.Ports, c.Status)
			var cont containerData
			cont.Id = c.ID
			cont.Image = c.Image
			cont.Status = c.Status
			cont.Names = strings.Join(c.Names, ",")
			cont.Labels = strings.Join(MapToArray(c.Labels), ",")
			cont.Ports = c.Ports
			contBulk.Host = HostForContainers
			contBulk.ContData = contBulk.addContainerData(cont)
		}
		if len(contBulk.ContData) > 0 {
			ch <- contBulk
		}
		log.Debug("Container monitor is going to sleep under the next collection")
		time.Sleep(time.Duration(client.contListIntervalSec) * time.Second)
	}
}
