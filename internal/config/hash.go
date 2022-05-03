package config

// HashConf holds the application's hash configurations.
type HashConf struct {
	MinLength int    `yaml:"min_length"`
	Alphabet  string `yaml:"alphabet"`
	Salt      string `yaml:"salt"`
}
