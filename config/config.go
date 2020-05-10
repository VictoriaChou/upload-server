package config

import (
	"errors"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	PG PG
}
type PG struct {
	Host     string
	Port     string
	User     string
	Password string
	UserName string
	FileName string
}

// global config item
var Cfg Config

func Parse(fileName string) error {
	if FileExists(fileName) {
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Panic("Database init error!")
			return err
		}
		_, err := toml.Decode(string(content), &Cfg)
		if err != nil {
			log.Panic("Config parse error!")
			return err
		}
		return nil
	}

	err := errors.New("File does not exists or is a directory")
}
