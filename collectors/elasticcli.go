package collectors

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	elastic "gopkg.in/olivere/elastic.v3"
)

var indexName = "dockermon"

type esClient struct {
	client *elastic.Client
}

// Create the Elasticsearch client to publish all the data collected
func NewEsClient(url string) esClient {
	log.Infof("Going to create Elasticsearch client with %v", url)
	client, err := elastic.NewSimpleClient(elastic.SetURL(url))
	if err != nil {
		log.Error("Failed to create Elasticsearch client", err)
		panic(err)
	}

	return esClient{client}
}

// Send container monitor data to Elasticsearch
func ReadAndSendContainerData(cli esClient, ch chan ContainersBulkData) {

	for {

		data := <-ch

		log.Debug("Going to send container data to ES: ", data)
		bdata, _ := json.Marshal(data)

		_, err := cli.client.Index().Index(indexName).Type(fetchDataType(bdata)).BodyString(string(bdata)).Do()
		if err != nil {
			log.Error("Problem when sending container data: ", err)
		}
	}

}

// Send image monitor data to Elasticsearch
func ReadAndSendImageData(cli esClient, ch chan ImageBulkData) {

	for {

		data := <-ch
		log.Debug("Going to send image data to ES: ", data)
		bdata, _ := json.Marshal(data)

		_, err := cli.client.Index().Index(indexName).Type(fetchDataType(bdata)).BodyString(string(bdata)).Do()
		if err != nil {
			log.Error("Problem when sending container data: ", err)
		}
	}
}

// Send docker events to Elasticsearch
func SendEvent(cli esClient, ch chan DockerEvent) {

	for {

		data := <-ch

		log.Debug("Going to send event to ES: ", data)

		eventstr, _ := json.Marshal(data)

		cli.client.Index().Index(indexName).Type(data.DataType).BodyString(string(eventstr)).Do()
	}

}

func fetchDataType(bdata []byte) string {

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(bdata), &dat); err != nil {
		panic(err)
	}

	return dat["DataType"].(string)

}
