package config

type Config struct {
	BindAddr string  `toml:"bind_addr"`
	Database string  `toml:"db_url"`
	TGToken  string  `toml:"tg_token"`
	Rec      RecData `toml:"rec_data"`
}

type RecData struct {
	GPTData GPTData `toml:"gpt_data"`
}

type GPTData struct {
	GPTToken    string  `toml:"gpt_token"`
	Model       string  `toml:"model"`
	Temperature float64 `toml:"temperature"`
	MaxTokens   int     `toml:"max_tokens"`
	Prompt      string  `toml:"prompt"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
