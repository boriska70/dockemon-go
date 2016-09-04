package collectors

import (
	elastic "gopkg.in/olivere/elastic.v3"
	log "github.com/Sirupsen/logrus"
        "encoding/json"
	"fmt"
)

var indexName = "dockermon"

type esClient struct {
	client *elastic.Client
}


func NewEsClient(url string) esClient {
	log.Infof("Going to create Elasticsearch client with %v", url)
	client, err := elastic.NewSimpleClient(elastic.SetURL(url))
	if err != nil {
		log.Error("Failed to create Elasticsearch client", err)
		panic(err)
	}

	return esClient{client}
}

func ReadAndSend(cli esClient, ch chan ContainersBulkData) {

	for {

		data := <-ch

		log.Info("Received %v", data)
		bdata, _ := json.Marshal(data)
		log.Info("Received json %s", string(bdata))
		fmt.Println(string(bdata))

		//cli.client.Index().Index(indexName).Type("extractfromdata").BodyString(data).Do()
	}

}