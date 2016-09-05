package collectors

import (
	"github.com/docker/distribution/context"
	"github.com/docker/engine-api/types"
	log "github.com/Sirupsen/logrus"
	"fmt"
_	"encoding/json"
)

func EventsCollect(client doClient) {

	options := types.EventsOptions{}
	b1 := make([]byte, 1024)

	for {

		body, err := client.dc.Events(context.Background(), options)

		if err != nil {
			log.Error(err)
		}

		n1, err := body.Read(b1)



		fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

//		myjson,_ := json.Marshal(b1[:n1])
		//fmt.Printf("%d bytes in JSON: %s\n", len(myjson), string(myjson))


		log.Info("buffer: ", len(b1))
		log.Info("body: ", n1)

//		time.Sleep(1 * time.Second)
		types.

	}
}
