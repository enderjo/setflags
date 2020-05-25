package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//Server 服务器配置
type Server struct {
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Db   DB     `mapstructure:"db" json:"db" yaml:"db"`
	App  App    `mapstructure:"app" json:"app" yaml:"app"`
}

//DB 数据库配置
type DB struct {
	DbType   string `mapstructure:"dbType" json:"dbType" yaml:"dbType"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	DbName   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

//App 应用配置
type App struct {
	Keystore string `mapstructure:"keystore" json:"keystore" yaml:"keystore"`
}

//AppKeystore 应用keystore
type AppKeystore struct {
	Pin        string `mapstructure:"pin" json:"pin" yaml:"pin"`
	ClientID   string `mapstructure:"client_id" json:"client_id" yaml:"client_id"`
	SessionID  string `mapstructure:"session_id" json:"session_id" yaml:"session_id"`
	PinToken   string `mapstructure:"pin_token" json:"pin_token" yaml:"pin_token"`
	PrivateKey string `mapstructure:"private_key" json:"private_key" yaml:"private_key"`
}

var (
	//ServerConfig 服务器配置
	ServerConfig Server
	//AppConfig 应用配置文件
	AppConfig AppKeystore

	//ServerViper 配置文件
	ServerViper *viper.Viper
	//AppViper 配置文件
	AppViper *viper.Viper
)

const defaultConfigFile = "conf/config.yaml"

func init() {
	var err error
	ServerViper, err = initServerConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error server config file: %s ", err))
	}

	AppViper, err = initAppConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error app keystore file: %s ", err))
	}
}

func initServerConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&ServerConfig); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&ServerConfig); err != nil {
		fmt.Println(err)
	}

	return v, err
}

func initAppConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(ServerConfig.App.Keystore)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("app keystore file changed:", e.Name)
		if err := v.Unmarshal(&AppConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&AppConfig); err != nil {
		fmt.Println(err)
	}

	return v, err
}
