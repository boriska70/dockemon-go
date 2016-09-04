package collectors

import (
	elastic "gopkg.in/olivere/elastic.v3"
	log "github.com/Sirupsen/logrus"
        "encoding/json"
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

//		log.Info("Received data: ", data)
		bdata, _ := json.Marshal(data)
//		log.Info("Received json: ", string(bdata))

		cli.client.Index().Index(indexName).Type(fetchDataType(bdata)).BodyString(string(bdata)).Do()
	}

}

func fetchDataType(bdata []byte) string {

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(bdata), &dat); err != nil {
		panic(err)
	}

	return dat["DataType"].(string)

}
