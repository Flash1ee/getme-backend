package internal

type RepositoryConnections struct {
	DataBaseUrl string `toml:"database_url"`
}

type Config struct {
	LogLevel   string                `toml:"log_level"`
	LogAddr    string                `toml:"log_path"`
	Domain     string                `toml:"domain"`
	BindAddr   string                `toml:"bind_addr"`
	Repository RepositoryConnections `toml:"repository"`
}
