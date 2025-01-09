package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

func Init() (err error) {
	// 设置配置文件名称（不带扩展名）
	viper.SetConfigName("config")
	// 设置配置文件路径（可以是相对路径或绝对路径）
	viper.AddConfigPath(".")
	// 设置配置文件类型
	viper.SetConfigType("yaml")

	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Fatal error config file: %w \n", err)
		return err
	}

	// 动态加载配置信息——修改后自动生效（可注释关闭）
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	return nil
}
