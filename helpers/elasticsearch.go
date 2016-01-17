package helpers

import (
	"gopkg.in/olivere/elastic.v3"
	"log"
	"os"
	"time"
)

var elasticsearch, _ = prepareElasticsearch()

func prepareElasticsearch() (*elastic.Client, error) {
	errorlog := log.New(os.Stdout, "Wave ", log.LstdFlags)
	client, err := elastic.NewClient(elastic.SetErrorLog(errorlog))
	if err != nil {
		log.Println(err)
		return client, err
	}
	exists, err := client.IndexExists("frames").Do()
	if err != nil {
		log.Println(err)
		return client, err
	}
	if exists {
		client.DeleteIndex("frames").Do()
	}
	return client, nil
}

func ElasticacheFrame(frame []byte) {
	record, err := elasticsearch.Index().
		Index("frames").
		Type("frame").
		BodyString(string(frame)).
		Do()
	if err != nil {
		log.Println(err)
	}
	go func() {
		time.Sleep(30 * time.Second)
		elasticsearch.Delete().
			Index("frames").
			Type("frame").
			Id(record.Id).
			Do()
	}()
}
