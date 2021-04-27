package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var GloCfg = &viper.Viper{}

func GCfg(path string) error { //配置文件
	GloCfg = viper.New()
	//设置配置文件的名字
	GloCfg.SetConfigName("dev")
	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	GloCfg.AddConfigPath(path)
	//设置配置文件类型
	GloCfg.SetConfigType("yaml")
	if err := GloCfg.ReadInConfig(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

