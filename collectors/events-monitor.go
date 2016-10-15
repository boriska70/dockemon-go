package collectors

import (
	"github.com/docker/docker/api/types"
	log "github.com/Sirupsen/logrus"
	"time"
	"context"
	"io"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/events"
)

type DockerEvent struct {
	Id             string
	Status         string
	Action         string
	From           string
	Type           string
	Actor          Actor
	CollectionTime time.Time
	DataType string
}

type Actor struct {
	ID         string
	Attributes map[string]string
}

func EventsCollect(client doClient, ch chan DockerEvent) {

	filters := filters.NewArgs()
	filters.Add("container-events", events.ContainerEventType)
	filters.Add("image-events", events.ImageEventType)
	options := types.EventsOptions{
		Filters: filters,
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
				log.Info(e.ID, e.Action, e.Status, e.Type)
			var de DockerEvent
				de.DataType = "DockerEvent"
				de.CollectionTime = time.Now()
				de.Type = e.Type
				de.Action = e.Action
				de.Status = e.Status
				de.Id = e.ID
				de.From = e.From
				de.Actor.ID = e.Actor.ID
			de.Actor.Attributes = e.Actor.Attributes
				if len(de.Action) > 0 {
					ch <-de
				}
			}
		}

	}
}