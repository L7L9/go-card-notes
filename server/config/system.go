package config

type SystemConfig struct {
	Addr          int    `mapstructure:"addr"`
	AdminAccount  string `mapstructure:"admin-account"`
	AdminPassword string `mapstructure:"admin-password"`
	RouterPrefix  string `mapstructure:"router-prefix"`
}
