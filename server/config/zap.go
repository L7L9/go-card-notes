package config

// ZapConfig //
// zap日志库的配置
type ZapConfig struct {
	Level        string `mapstructure:"level"`
	Director     string `mapstructure:"director"`
	LogInConsole bool   `mapstructure:"log-in-console"`
}
