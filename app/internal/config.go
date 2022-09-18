package internal

type RepositoryConnections struct {
	DataBaseUrl      string `toml:"database_url"`
	DataBaseUrlLocal string `toml:"database_url_local"`
	MySQLDatabaseUrl string `toml:"database_url_mysql"`
	CacheURL         string `toml:"cache_url"`
	CacheURLLocal    string `toml:"cache_url_local"`
}
type Microservices struct {
	SessionServerUrl      string `toml:"session_url"`
	SessionServerUrlLocal string `toml:"session_url_local"`

	SessionRedisUrl      string `toml:"session_redis_url"`
	SessionRedisUrlLocal string `toml:"session_redis_url_local"`
}
type TelegramAuth struct {
	Token string `toml:"token"`
}

type Config struct {
	LogLevel      string                `toml:"log_level"`
	LogAddr       string                `toml:"log_path"`
	Domain        string                `toml:"domain"`
	BindAddr      string                `toml:"bind_addr"`
	Repository    RepositoryConnections `toml:"repository"`
	Microservices Microservices         `toml:"microservices"`
	TgAuth        TelegramAuth          `toml:"telegram"`
	DebugMode     bool                  `toml:"debug"`
}
