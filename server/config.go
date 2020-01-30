package server

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type config struct {
	UploadDir string `yaml:"upload_dir"`
	Domain    string `yaml:"domain"`
}

func parseConfig() *config {
	conf := &config{}
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return conf
}
