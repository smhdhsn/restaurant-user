package config

// ServerConf holds the application's server configurations.
type ServerConf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
