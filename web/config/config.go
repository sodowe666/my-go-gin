package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

type conf struct {
	Server struct {
		Name        string `yaml:"name"`
		Port        int    `yaml:"port"`
		Environment string `yaml:"environment"`
		Session     struct {
			SessionId string `yaml:"sessionId"`
		} `yaml:"session"`
		Identity struct {
			IdentityCookie struct {
				Name     string `yaml:"name"`
				HttpOnly string `yaml:"httpOnly"`
			} `yaml:"identityCookie"`
		} `yaml:"identity"`
	}
	Jwt struct {
		Name     string `yaml:"name"`
		Secret   string `yaml:"secret"`
		Duration int    `yaml:"duration"`
	} `yaml:"jwt"`
	Db struct {
		DBType       string `yaml:"type"`
		DataBase     string `yaml:"database"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		Charset      string `yaml:"charset"`
		Host         string `yaml:"host"`
		Port         string `yaml:"port"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
		IdleTime     int    `yaml:"idleTime"`
		Slaves       []struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
		} `yaml:"slaves"`
	} `yaml:"db"`
	Redis struct {
		Host         string `yaml:"host"`
		Port         string `yaml:"port"`
		Password     string `yaml:"password"`
		Database     int    `yaml:"database"`
		IdleTime     int    `yaml:"idleTime"`
		MaxActive    int    `yaml:"maxActive"`
		ReadTimeout  int    `yaml:"readTimeout"`
		WriteTimeout int    `yaml:"writeTimeout"`
	} `yaml:"redis"`
	ElasticSearch struct {
		Addr                string `yaml:"addr"`
		User                string `yaml:"user"`
		Password            string `yaml:"password"`
		HealthCheckInterval int    `yaml:"healthCheckInterval"`
	} `yaml:"elasticSearch"`
}

var cfg *conf
var once sync.Once

func init() {
	loadYml()
}

func GetConfig() *conf {
	if cfg == nil {
		loadYml()
	}
	return cfg
}

/**
加载yml文件
*/
func loadYml() *conf {
	once.Do(func() {
		if cfg == nil { //指针指向nil，不存在的地址
			cfg = new(conf) //初始化conf，给地址给cfg指针
			ymlFile, err := ioutil.ReadFile("/Users/jqz/go/src/web/config/config.yml")
			if err != nil {
				panic(err)
			}
			err = yaml.Unmarshal(ymlFile, cfg)
			if err != nil {
				panic(err)
			}
		}
	})
	return cfg
}
