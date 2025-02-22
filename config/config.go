package config

import (
	"strings"

	"github.com/spf13/viper"
)

var Cfg Config

type KeyBoardConfig struct {
	Keys     []string `yaml:"keys"`
	DoneKeys []string `yaml:"donekeys"`
}

type Detail struct {
	Api string `yaml:"api"`
	Url string `yaml:"url"`
}

type Config struct {
	KeyBoard KeyBoardConfig    `yaml:"keyboard"`
	Model    map[string]Detail `yaml:"model"`
	Server   Server            `yaml"json:"server"`
}

type Server struct {
	Port string `yaml:"port"`
}

func InitConfig() {
	// wd, _ := os.Getwd()
	wd := "E:\\screenshot-ai" // test
	// fmt.Println("wd" + wd)
	newWd := strings.Replace(wd, "\\", "/", 2)
	configPath := newWd + "/config/config.yaml"

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("file not found")
		} else {
			panic("fail to read config file")
		}
	}
	// fmt.Println("读取的原始配置:", viper.AllSettings())
	err := viper.Unmarshal(&Cfg)
	// fmt.Println(Cfg)
	if err != nil {
		panic("read config error")
	}
}
