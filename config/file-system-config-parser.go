package config

import (
	"errors"
	"os"

	util "github.com/dimitrijed93/load-balancer/util"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type FileSystemConfigParser struct {
	data LoadBalancerConfig
}

func (ycp FileSystemConfigParser) Parse(input string) (*LoadBalancerConfig, error) {
	if input == util.EMPTY_STRING {
		return &LoadBalancerConfig{}, errors.New("Input file is empty")
	}
	configFile, err := os.ReadFile(input)
	if err != nil {
		return &LoadBalancerConfig{}, errors.New("Unable to read config file")
	}
	var config LoadBalancerConfig
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return &LoadBalancerConfig{}, err
	}
	log.Info().Msgf("FileSystemConfigParser >> Loaded config %s", config)
	return &config, nil
}
