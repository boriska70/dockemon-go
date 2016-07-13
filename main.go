package main


import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/boriska70/dockermon-go/collectors"
	"flag"
)

var timeLayout="2006-01-02 15:04:05";

func init() {
	log.SetLevel(log.DebugLevel);
}

func main() {

	cci := flag.Int64("cci", 60, "Container Collection Interval")
	log.Debugf("Running with container collection interval %v", *cci)

	log.Infof("Start running dockermon-go at %v", time.Now().Format(timeLayout));

	client := collectors.NewClient(*cci)
	log.Infof("Client: %v", client)
	go client.ListContainers();
	go client.Cl.StartMonitorEvents(client.EventCallBack, nil);
	time.Sleep(3600 * time.Second)



}