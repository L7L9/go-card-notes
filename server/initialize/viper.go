package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"lqlzzz/go-card-notes/global"
)

// InitViper //
// initialize viper config
func InitViper() *viper.Viper {
	v := viper.New()

	// 设置配置文件名字和类型
	v.SetConfigName("config.yaml")
	v.SetConfigType("yaml")
	err := v.ReadInConfig()

	// 读取配置文件失败
	if err != nil {
		// 失败则程序直接中断
		panic(fmt.Errorf("failed to upload config file: %s \n", err))
	}

	// 将配置文件中的数据解析到结构体中
	if err = v.Unmarshal(&global.GCN_CONFIG); err != nil {
		// 失败则程序直接中断
		panic(fmt.Errorf("failed to upload config file: %s \n", err))
	}
	return v
}
