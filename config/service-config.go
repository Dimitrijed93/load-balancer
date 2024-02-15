package config

import (
	"strings"

	util "github.com/dimitrijed93/load-balancer/util"
)

type ServiceConfig struct {
	Name              string             `yaml:"name"`
	Type              string             `yaml:"type"`
	Path              string             `yaml:"path"`
	ServersConfig     []ServerConfig     `yaml:"servers"`
	RateLimiterConfig *RateLimiterConfig `yaml:"rate-limiter"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func NewServiceConfig(uri string) ServiceConfig {

	cnf := NewLoadBalancerConfig()

	var svc ServiceConfig
	for _, item := range cnf.ServicesConfig {
		if strings.HasPrefix(uri, item.Path) {
			svc = item
		}
	}
	if svc.Path == util.EMPTY_STRING {
		panic("ServiceConfig >> Invalid path, no configuration found for path " + uri)
	}
	return svc
}
