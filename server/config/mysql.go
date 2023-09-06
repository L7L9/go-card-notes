package config

// MysqlConfig //
// 数据库配置
type MysqlConfig struct {
	Port        string `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	Host        string `mapstructure:"host"`
	DbName      string `mapstructure:"db-name"`
	MaxOpenConn int    `mapstructure:"max-open-conn"`
	MaxIdleConn int    `mapstructure:"max-idle-conn"`
}

// Dsn //
// 获取数据库dsn
func (mysql *MysqlConfig) Dsn() string {
	return mysql.Username + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/" + mysql.DbName + "?charset=utf8&parseTime=true"
}
