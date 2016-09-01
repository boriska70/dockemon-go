package collectors

import (
	elastic "gopkg.in/olivere/elastic.v3"
	log "github.com/Sirupsen/logrus"
)

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