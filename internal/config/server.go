package config

// ServerConf holds the application's server configurations.
type ServerConf struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}
