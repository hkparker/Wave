package helpers

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/olivere/elastic.v3"
	"time"
)

var elasticsearch *elastic.Client

func SetupElasticsearch() {
	client, err := elastic.NewClient() // pass logrus logger
	if err != nil {
		log.Println(err)
		elasticsearch = client
		return
	}
	exists, err := client.IndexExists("frames").Do()
	if err != nil {
		log.Println(err)
		elasticsearch = client
		return
	}
	if exists {
		client.DeleteIndex("frames").Do()
	}
	elasticsearch = client
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
	go func() { // TTL depicated in ES2.2 but still works?
		time.Sleep(30 * time.Second)
		elasticsearch.Delete().
			Index("frames").
			Type("frame").
			Id(record.Id).
			Do()
	}()
}
