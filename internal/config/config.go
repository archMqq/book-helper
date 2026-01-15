package config

type Config struct {
	BindAddr string `toml:"bind_addr"`
	Database string `toml:"db_url"`
	TGToken  string `toml:"tg_token"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
