package core

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"os"
	"sync"
	"time"
	"web/config"
)

type esClient struct {
	once   sync.Once
	client *elastic.Client
}
var es esClient

func init() {
	initElastic()
}

func Elastic() *elastic.Client {
	if es.client  == nil {
		initElastic()
	}
	return es.client
}

func initElastic() {
	es.once.Do(func() {
		if es.client == nil {
			cfg := config.GetConfig()
			elasticCfg := cfg.ElasticSearch
			err := os.MkdirAll("./elastic", 0777)
			if err != nil {
				panic(err.Error())
			}
			f, err := os.Create("./elastic/error.log")
			if err != nil {
				panic("elastic error.log create fail. " + string(err.Error()))
			}
			es.client, err = elastic.NewClient(
				elastic.SetURL(elasticCfg.Addr),
				elastic.SetBasicAuth(elasticCfg.User, elasticCfg.Password),
				elastic.SetHealthcheckInterval(time.Duration(elasticCfg.HealthCheckInterval)*time.Second),
				elastic.SetErrorLog(log.New(f, "ERROR ", log.LstdFlags)),
			)
			if err != nil {
				panic(err.Error())
			}
			info, code, err := es.client.Ping(elasticCfg.Addr).Do(context.Background())
			if err != nil {
				panic("elastic Ping error")
			}
			fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
		}
	})
}
