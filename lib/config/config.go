package config

type httpConfig struct {
	HttpServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
}
type dbConfig struct {
	DbDriver string `mapstructure:"DB_DRIVER"`
	DbSource string `mapstructure:"DB_SOURCE"`
}

type Config struct {
	HTTP httpConfig
	DB   dbConfig
}
