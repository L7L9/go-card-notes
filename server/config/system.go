package config

// SystemConfig //
// 本系统配置
type SystemConfig struct {
	Addr          int    `mapstructure:"addr"`
	AdminAccount  string `mapstructure:"admin-account"`
	AdminPassword string `mapstructure:"admin-password"`
}
