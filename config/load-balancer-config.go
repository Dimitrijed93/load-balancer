package config

import (
	"os"

	util "github.com/dimitrijed93/load-balancer/util"
	"github.com/rs/zerolog/log"
)

var (
	lbConfig *LoadBalancerConfig
)

type LoadBalancerConfig struct {
	ServicesConfig []ServiceConfig `yaml:"services"`
}

func NewLoadBalancerConfig() *LoadBalancerConfig {
	if lbConfig == nil {
		cp, err := NewConfigParser(os.Getenv(util.ENV_VAR_CONFIG_PATH))
		cnf := cp.Config
		if err != nil {
			log.Error().Msg("Unable to open config")
		}
		lbConfig = &cnf
		return &cnf
	}
	return lbConfig
}
