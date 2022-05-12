package internal

type RepositoryConnections struct {
	DataBaseUrl string `toml:"database_url"`
}
type Microservices struct {
	SessionServerUrl string `toml:"session_url"`
	SessionRedisUrl  string `toml:"session_redis_url"`
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
}
