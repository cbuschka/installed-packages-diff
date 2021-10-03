package config

type ServerConfig struct {
	Hostname string
	Port     int
	Username string
}

type ServersConfig struct {
	Servers []ServerConfig
}

type Config struct {
	Groups map[string]ServersConfig
}
