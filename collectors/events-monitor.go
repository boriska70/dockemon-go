package collectors

import (
	"github.com/docker/distribution/context"
	"github.com/docker/engine-api/types"
	log "github.com/Sirupsen/logrus"
	"fmt"
	"time"
	"encoding/json"
)

type DockerEventPackage struct {
	event dockerEvent
	DataType string
	CollectionTime time.Time
}

type dockerEvent struct {
	id string
	status string
	from string
	sourceType string
	action string
	actor actor
	eventTime time.Time
}

type actor struct {
	actorId string
	attributes [] string
}

type AAA struct {
	DockEvent string
	DataType string
	CollectionTime time.Time
}


func EventsCollect(client doClient, ch chan []byte) {

	options := types.EventsOptions{}

	for {
		b1 := make([]byte, 1024)
		body, err := client.dc.Events(context.Background(), options)

		if err != nil {
			log.Error(err)
		}

		var event string

		n1, err := body.Read(b1)
		fmt.Println("Body length is ", n1)

		json.NewDecoder(body).Decode(event)
		fmt.Println("Event is ", string(b1[:n1]))

/*		var aaa AAA
		aaa.event =string(b1[1:n1-1])
		aaa.CollectionTime = time.Now()
		aaa.DataType="DockerEvent"*/


				//if len(aaa.event) > 0 {
					if len(b1) > 0 {
					ch <- b1[:n1]
				}


	}
}
