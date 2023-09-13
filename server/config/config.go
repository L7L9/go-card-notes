package config

type Config struct {
	Mysql  MysqlConfig  `mapstructure:"mysql"`
	Zap    ZapConfig    `mapstructure:"zap"`
	Jwt    JwtConfig    `mapstructure:"jwt"`
	System SystemConfig `mapstructure:"system"`
}
