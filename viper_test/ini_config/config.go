package ini_config

import (
	"os"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var (
	once sync.Once
)

var config *Config

func InitViper(filePath string) {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("ini")
		viper.AddConfigPath(filePath)

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				logrus.Fatal("配置文件未找到")
				os.Exit(1)
			}
		}

		includes := viper.GetStringSlice("includes")

		for _, file := range includes {
			viper.SetConfigName(file)
			viper.SetConfigType("ini")
			err := viper.MergeInConfig()
			if err != nil {
				logrus.Fatal("配置文件读取出错")
			}
		}
		viper.SetEnvPrefix("ACS")
		viper.AutomaticEnv()
		replacer := strings.NewReplacer(".", "__")
		viper.SetEnvKeyReplacer(replacer)
		// 监控配置和重新获取配置
		viper.WatchConfig()

		config = new(Config)
		err := viper.Unmarshal(config)
		if err != nil {
			logrus.Error(err)
			os.Exit(1)
		}
		logrus.Printf("config: %+v\n", config)
	})
}
