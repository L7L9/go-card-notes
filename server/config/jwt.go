package config

type JwtConfig struct {
	SigningKey  string `mapstructure:"signing-key"`
	Issuer      string `mapstructure:"issuer"`
	ExpiresTime int    `mapstructure:"expires-time"`
}
