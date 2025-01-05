package config

type HttpConfig struct {
	HttpServerAddress string
}

type DbConfig struct {
	DbDriver string
	DbSource string
}

type Config struct {
	HttpServerAddress string `json:"HTTP_SERVER_ADDRESS"`
	DbDriver          string `json:"DB_DRIVER"`
	DbSource          string `json:"DB_SOURCE"`
}
