package conf

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func RunViper() {
	absolutePath, err := filepath.Abs("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	viper.SetConfigFile(absolutePath)
	viper.ReadInConfig()

	// 注册每次配置文件发生变更后都会调用的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config file changed: %s\n", e.Name)
	})

	// 监控并重新读取配置文件，需要确保在调用前添加了所有的配置路径
	viper.WatchConfig()

	// 阻塞程序，这个过程中可以手动去修改配置文件内容，观察程序输出变化
	time.Sleep(time.Second * 10)

	// 读取配置值
	fmt.Printf("username: %s\n", viper.Get("username"))
}
