package config

type LoadBalancerConfig struct {
	ServicesConfig []ServiceConfig `yaml:"services"`
}

type ServiceConfig struct {
	Name              string            `yaml:"name"`
	Type              string            `yaml:"type"`
	Path              string            `yaml:"path"`
	ServersConfig     []ServerConfig    `yaml:"servers"`
	RateLimiterConfig RateLimiterConfig `yaml:"rate-limiter"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type RateLimiterConfig struct {
	Name        string `yaml:"name"`
	MaxRequests uint32 `yaml:"max-requests"`
	WindowSize  uint32 `yaml:"window-size"`
}

func (rl RateLimiterConfig) IsEmpty() bool {
	return rl.MaxRequests == 0 && rl.WindowSize == 0
}
