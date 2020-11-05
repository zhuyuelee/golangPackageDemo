package utils

import (
	config "GoSql/EchoDemo/dtos/Config"
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const configPath = "config.yaml"

var configData string

func init() {
	data, err := ioutil.ReadFile(configPath)
	if err != nil && err != io.EOF {
		conf := config.Config{
			CORS: config.CORS{
				Domains: []string{"http://127.0.0.1:9090"},
			},
			DataBase: config.DataBase{
				Sqlite: "data/data.db",
			},
		}
		data, err = yaml.Marshal(&conf)
		if err == nil {
			err = ioutil.WriteFile(configPath, data, os.ModePerm)
		}
	}
	if err == nil {
		configData = string(data)
	}
}

//GetCORSConfig 获取CORS配置
func GetCORSConfig() (domains []string, err error) {

	if configData != "" {
		cors := &config.CORS{}
		err = yaml.Unmarshal([]byte(configData), cors)
		if err == nil {
			domains = cors.Domains
		}
	}
	return
}

//GetDatabaseConfig 获取数据库配置
func GetDatabaseConfig() (database string, err error) {

	if configData != "" {
		conf := &config.DataBase{}
		err = yaml.Unmarshal([]byte(configData), conf)
		if err == nil {
			database = conf.Sqlite
		}
	}
	return
}
