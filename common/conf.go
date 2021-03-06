package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var GlobalConfig Config

type Config struct {
	OSS OSS `yaml:"OSS"`
	DB  DB  `yaml:"DB"`
}

type OSS struct {
	Bucket           string `yaml:"Bucket"`
	Endpoint         string `yaml:"Endpoint"`         //外网 Endpoint
	EndpointInternal string `yaml:"EndpointInternal"` //内网 Endpoint
	AccessKey        string `yaml:"AccessKey"`
	SecretKey        string `yaml:"SecretKey"`
	UseInternal      bool   `yaml:"UseInternal"` //是否使用内网 Endpoint
	UseSSL           bool   `yaml:"UseSSL"`      //是否使用SSL(对于minio有用)
}

type DB struct {
	DSN     string `yaml:"DSN"`
	MaxIdle int    `yaml:"MaxIdle"`
	MaxOpen int    `yaml:"MaxOpen"`
}

func InitConfig(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(contents, &GlobalConfig)
	if err != nil {
		return err
	}
	return nil
}
