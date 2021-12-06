package core

import (
	"db-core/helpers"
	"github.com/goccy/go-yaml"
	"io/ioutil"
	"log"
	"os"
)

type DBConfig struct {
	DSN         string `yaml:"dsn"`
	MaxOpenConn int    `yaml:"maxOpenConn"`
	MaxIdleConn int    `yaml:"maxIdleConn"`
	MaxLifeTime int    `yaml:"maxLifeTime"`
}

type SysConfigStruct struct {
	DBConfig *DBConfig `yaml:"dbConfig"`
	ApiList  []*API    `yaml:"apis"`
}

func (c *SysConfigStruct) FindAPI(name string) *API {
	for _, api := range c.ApiList {
		if api.Name == name {
			return api
		}
	}
	return nil
}

var SysConfig *SysConfigStruct

const SysConfigPath = "./app.yml"

func InitConfig() {
	config := &SysConfigStruct{}
	f, err := os.Open(SysConfigPath)
	helpers.Error(err, "找不到配置文件")

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal([]byte(data), config)
	helpers.Error(err, "配置文件不正确")
	SysConfig = config
}
