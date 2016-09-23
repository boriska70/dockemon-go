package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/boriska70/dockermon-go/collectors"
	"flag"
)

var timeLayout = "2006-01-02 15:04:05";

func init() {
	log.SetLevel(log.InfoLevel);
}

func main() {
	cci := flag.Int64("cci", 60, "Container Collection Interval")
	esurl := flag.String("esurl", "http://localhost:9200", "Elasticsearch URL")
	flag.Parse()
	log.Debugf("Running with container collection interval %v seconds", *cci)
	log.Debugf("Collected data will be sent to %v", *esurl)

	log.Infof("Start running dockermon-go at %v", time.Now().Format(timeLayout));

	client := collectors.NewDockerClient(*cci)
	collectors.SetHostStaticData(client)
	log.Infof("Docker client created for %v", collectors.Host)
	elasticclient := collectors.NewEsClient(*esurl);
	log.Infof("Elastic? ", elasticclient)
	contChannel := make(chan collectors.ContainersBulkData)
	eventChannel := make(chan collectors.DockerEvent)
	imageChannel := make(chan collectors.ImageBulkData)

	go collectors.ReadAndSendImageData(elasticclient, imageChannel)
	go collectors.ReadAndSendContainerData(elasticclient, contChannel)
	go collectors.SendEvent(elasticclient, eventChannel)

	go collectors.ImageStats(client, imageChannel)
	go collectors.ContainerStats(client, contChannel)
	go collectors.EventsCollect(client, eventChannel)

	//go client.Cl.StartMonitorEvents(client.EventCallBack, nil);
	for {
		time.Sleep(3600 * time.Second)
	}

}