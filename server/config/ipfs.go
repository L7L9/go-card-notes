package config

// IpfsConfig //
// ipfs配置
type IpfsConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"post"`
}
