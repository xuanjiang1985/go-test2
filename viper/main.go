package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml") // 指定配置文件
	viper.AddConfigPath("./")          // 指定查找配置文件的路径
	err := viper.ReadInConfig()        // 读取配置信息
	if err != nil {                    // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监控配置文件变化
	viper.WatchConfig()

	fmt.Printf("---type: %T ---value: %v\n", viper.Get("version"), viper.Get("version"))

	fmt.Printf("---type: %T ---value: %v\n", viper.Get("user.sex"), viper.Get("user.sex"))

	fmt.Printf("---type: %T ---value: %v\n", viper.Get("database"), viper.Get("database"))

	fmt.Printf("---type: %T ---value: %v\n", viper.Get("etcd"), viper.Get("etcd"))

}
