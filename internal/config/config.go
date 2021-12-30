package config

type ServerConfig struct {
	Url  string `yaml:"url"`
	Type string `yaml:"type"`
}

type ServersConfig struct {
	Servers []ServerConfig
}

type Config struct {
	Version string                   `yaml:"version"`
	Groups  map[string]ServersConfig `yaml:"groups"`
}
