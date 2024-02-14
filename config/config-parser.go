package config

type LoadBalancerConfigParser interface {
	parse(input interface{}) (LoadBalancerConfig, error)
}
