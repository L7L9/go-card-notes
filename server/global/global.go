package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"lqlzzz/go-card-notes/config"
)

var (
	GCN_DB     *gorm.DB
	GCN_VIPER  *viper.Viper
	GCN_LOGGER *zap.Logger
	GCN_CONFIG config.Config
)
