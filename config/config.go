package config

import (
	"embed"
	build "github.com/Niexiawei/golang-utils/buildtools"
	"gopkg.in/yaml.v3"
	"io/fs"
)

var (
	Instance *Config
)

//go:embed *.yaml
var configFiles embed.FS

type Config struct {
	Mysql      Mysql      `yaml:"mysql"`
	HttpServer HttpServer `yaml:"http"`
	Redis      Redis      `yaml:"redis"`
}
type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       string `yaml:"database"`
	Pool     int    `yaml:"pool"`
}

type HttpServer struct {
	Port int `yaml:"port"`
}

type Redis struct {
	Pool     int    `yaml:"pool"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

func LoadConfig() {
	Instance = &Config{}
	if cByte, err := fs.ReadFile(configFiles, "config.yaml"); err == nil {
		_ = yaml.Unmarshal(cByte, Instance)
	} else {
		panic(err)
	}
	if build.IsProd() {
		if cByte, err := fs.ReadFile(configFiles, "config.prod.yaml"); err == nil {
			_ = yaml.Unmarshal(cByte, Instance)
		} else {
			panic(err)
		}
	}
}
