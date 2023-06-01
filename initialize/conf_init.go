package initialize

import (
	"blog_server/config"
	"blog_server/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// InitComnf 读取yaml配置文件 通过yaml.v2 这个工具
func InitComnf() {
	const configFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("get yamlCong error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlConf load Init success")
	global.Config = c
}
