package collectors

import (
	"github.com/docker/distribution/context"
	"github.com/docker/engine-api/types"
	log "github.com/Sirupsen/logrus"
	"time"
	"encoding/json"
	"bytes"
)

/*type DockerEventPackage struct {
	Event    *DockerEvent
	DataType string
}*/

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

	options := types.EventsOptions{}

	for {
//		var dep DockerEvent
		var de DockerEvent
		//		var str string

		b1 := make([]byte, 1024)
		body, err := client.dc.Events(context.Background(), options)

		if err != nil {
			log.Error(err)
		}

		n1, _ := body.Read(b1)
		log.Debug("Event body length is ", n1)

		//		str = string(b1[:n1])
		//		fmt.Println("String: ", str)
		dec := json.NewDecoder(bytes.NewReader(b1[:n1]))
		dec.Decode(&de)
		de.CollectionTime = time.Now()
		//		fmt.Println(de.Action)
		/*		destring, _ := json.Marshal(de)
		fmt.Println("Decoded: ", string(destring))
		fmt.Println("Done")*/

//		dep.Event = &de
		de.DataType = "DockerEvent"

		if len(de.Action) > 0 {
			if len(b1) > 0 {
				ch <- de
			}

		}
	}
}