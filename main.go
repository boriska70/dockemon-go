package main

import (
	"time"

	"flag"
	log "github.com/Sirupsen/logrus"
	"github.com/boriska70/dockermon-go/collectors"
)

var timeLayout = "2006-01-02 15:04:05"

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	dci := flag.Int64("dci", 60, "Data Collection Interval")
	esurl := flag.String("esurl", "http://localhost:9200", "Elasticsearch URL")
	flag.Parse()

	log.Infof("Start running dockermon-go at %v", time.Now().Format(timeLayout))
	defer log.Infof("Data collection is successfully started; data is sent to %v with interval %v seconds", esurl, dci)

	client := collectors.NewDockerClient(*dci)
	collectors.SetHostStaticData(client)
	log.Infof("Docker client created for %v", collectors.Host)

	elasticclient := collectors.NewEsClient(*esurl)
	log.Infof("Elastic? ", elasticclient)
	contChannel := make(chan collectors.ContainersBulkData)
	eventChannel := make(chan collectors.DockerEvent)
	imageChannel := make(chan collectors.ImageBulkData)

	//activate Elasticsearch client
	go collectors.ReadAndSendImageData(elasticclient, imageChannel)
	go collectors.ReadAndSendContainerData(elasticclient, contChannel)
	go collectors.SendEvent(elasticclient, eventChannel)

	//start data collection
	go collectors.ImageStats(client, imageChannel)
	go collectors.ContainerStats(client, contChannel)
	go collectors.EventsCollect(client, eventChannel)

	for {
		time.Sleep(3600 * time.Second)
	}

}
