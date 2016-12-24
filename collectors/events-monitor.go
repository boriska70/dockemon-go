package collectors

import (
	"context"
	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/api/types/filters"
	"io"
	"time"
)

// Standard docker event
type DockerEvent struct {
	Id             string
	Action         string
	From           string
	Type           string
	Actor          Actor
	CollectionTime time.Time
	DataType       string
}

// Actor parameter of the docker event
type Actor struct {
	ID         string
	Attributes map[string]string
}

// Catch image and container docker events and push them to data channel for further sending to storage by the relevant client
func EventsCollect(client doClient, ch chan DockerEvent) {

	eventFilters := filters.NewArgs()
	eventFilters.Add("container-events", events.ContainerEventType)
	eventFilters.Add("image-events", events.ImageEventType)
	options := types.EventsOptions{
		Filters: eventFilters,
	}

	for {

		//b1 := make([]byte, 1024)
		messages, errs := client.dc.Events(context.Background(), options)

	loop:
		for {
			select {
			case err := <-errs:
				if err != nil && err != io.EOF {
					log.Fatal(err)
				}

				break loop
			case e := <-messages:
				log.Infof("Event collected: %v, %v, %v", e.ID, e.Action, e.Type)
				var de DockerEvent
				de.DataType = "DockerEvent"
				de.CollectionTime = time.Now()
				de.Type = e.Type
				de.Action = e.Action
				de.Id = e.ID
				de.From = e.From
				de.Actor.ID = e.Actor.ID
				de.Actor.Attributes = e.Actor.Attributes
				if len(de.Action) > 0 {
					ch <- de
				}
			}
		}

	}
}
