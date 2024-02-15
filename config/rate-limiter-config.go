package config

type RateLimiterConfig struct {
	Name        string `yaml:"name"`
	MaxRequests uint32 `yaml:"max-requests"`
	WindowSize  int    `yaml:"window-size"`
}

func NewRateLimiterConfig(uri string) *RateLimiterConfig {
	svcConfig := NewServiceConfig(uri)
	if svcConfig.RateLimiterConfig != nil {
		return svcConfig.RateLimiterConfig
	}
	return &RateLimiterConfig{}
}

func (rl RateLimiterConfig) IsEmpty() bool {
	return rl.MaxRequests == 0 && rl.WindowSize == 0
}
