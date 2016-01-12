package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gopkg.in/olivere/elastic.v3"
	"log"
	"net/http"
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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// make sure the client presents proper TLS client cert
		return true
	},
}

func elasticache(frame []byte) {
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

func PollCollector(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err == nil {
		defer conn.Close()
		for {
			_, frame_bytes, err := conn.ReadMessage()
			if err != nil {
				break
			}
			//frame := string(frame_bytes)
			elasticache(frame_bytes)
			// update visualizer
		}
	}
}
