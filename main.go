package main


import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/boriska70/dockermon-go/collectors"
)

var timeLayout="2006-01-02 15:04:05";

func init() {
	log.SetLevel(log.DebugLevel);
}

func main() {

	log.Infof("Star running dockermon-go at %v", time.Now().Format(timeLayout));

	client := collectors.NewClient()
	log.Infof("Client: %v", client)
	client.ListContainers();
	client.Cl.StartMonitorEvents(client.EventCallBack, nil);
	time.Sleep(3600 * time.Second)



}