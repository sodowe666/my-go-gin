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

var elasticClient *elastic.Client
var esOnce sync.Once

func init() {
	initElastic()
}

func Elastic() *elastic.Client {
	if elasticClient == nil {
		initElastic()
	}
	return elasticClient
}

func initElastic() {
	esOnce.Do(func() {
		if elasticClient == nil {
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
			elasticClient, err = elastic.NewClient(
				elastic.SetURL(elasticCfg.Addr),
				elastic.SetBasicAuth(elasticCfg.User, elasticCfg.Password),
				elastic.SetHealthcheckInterval(time.Duration(elasticCfg.HealthCheckInterval)*time.Second),
				elastic.SetErrorLog(log.New(f, "ERROR ", log.LstdFlags)),
			)
			if err != nil {
				panic(err.Error())
			}
			info, code, err := elasticClient.Ping(elasticCfg.Addr).Do(context.Background())
			if err != nil {
				panic("elastic Ping error")
			}
			fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
		}
	})
}
