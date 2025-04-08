package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type config struct {
	Database database
	Redis    redis
	System   system
	Grpc     grpc
	Alisms   alisms
}

type database struct {
	DBPath string `mapstructure:"db_path"`
}

type redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

type system struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	Cert   string `mapstructure:"cert"` //证书
	Key    string `mapstructure:"key"`  //证书
	Static string `mapstructure:"static"`
}

type grpc struct {
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	RootCa    string `mapstructure:"root-ca"`
	ClientPem string `mapstructure:"client-pem"`
	ClientKey string `mapstructure:"client-key"`
	ServerPem string `mapstructure:"server-pem"`
	ServerKey string `mapstructure:"server-key"`
}

type alisms struct {
	RegionId        string `mapstructure:"regionId"`
	AccessKeyId     string `mapstructure:"accessKeyId"`
	AccessKeySecret string `mapstructure:"accessKeySecret"`
}

var c config

func Config() *config {
	return &c
}

func unmarshal() {
	if err := viper.Unmarshal(&c); err != nil {
		log.Panic("unable to decode into struct,", err)
	}
}

func abs(path string) string {
	if path == "" {
		_, filename, _, _ := runtime.Caller(0)
		return filepath.Dir(filename)
	} else {
		s, _ := filepath.Abs(path)
		if filepath.IsAbs(s) {
			return s
		} else {
			return ""
		}
	}
}

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(abs(""))
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Panic("Fatal error config file: ", err)
	}

	unmarshal()
	viper.OnConfigChange(func(e fsnotify.Event) {
		unmarshal()
		log.Println("Config file changed:", e.String())
	})
	viper.WatchConfig()
}
