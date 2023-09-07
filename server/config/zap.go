package config

type ZapConfig struct {
	Level        string `mapstructure:"level"`
	Director     string `mapstructure:"director"`
	LogInConsole bool   `mapstructure:"log-in-console"`
}
