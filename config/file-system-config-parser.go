package config

import (
	"errors"
	"os"

	util "github.com/dimitrijed93/load-balancer/util"
	"gopkg.in/yaml.v2"
)

type FileSystemConfigParser struct {
	Config LoadBalancerConfig
}

func NewConfigParser(input string) (*FileSystemConfigParser, error) {
	if input == util.EMPTY_STRING {
		return &FileSystemConfigParser{}, errors.New("input file is empty")
	}
	configFile, err := os.ReadFile(input)
	if err != nil {
		return &FileSystemConfigParser{}, errors.New("unable to read config file")
	}
	var config LoadBalancerConfig
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return &FileSystemConfigParser{}, err
	}
	return &FileSystemConfigParser{Config: config}, nil
}
