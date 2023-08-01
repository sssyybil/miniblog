package miniblog

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"miniblog/internal/pkg/log"
	"os"
	"path/filepath"
	"strings"
)

/**
helper.go 一般用来存放一些工具类的函数/方法
*/

const (

	// recommendedHomeDir 定义放置 miniblog 服务配置的默认目录
	recommendedHomeDir = ".miniblog"

	// defaultConfigName 指定了 miniblog 服务的默认配置文件名
	defaultConfigName = "miniblog.yaml"
)

func initConfig() {
	if cfgFile != "" {
		// 从命令行选项指定的配置文件中读取
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找用户主目录
		homeDir, err := os.UserHomeDir()

		// 如果用户主目录获取失败，打印 'Error: XXX' 错误，并退出程序（error code 1）
		cobra.CheckErr(err)

		fmt.Printf("homeDir: %v\n", recommendedHomeDir)

		// 将 `$HOME/<recommendedHomeDir>` 目录加入到配置文件的搜索路径中
		join := filepath.Join(homeDir, recommendedHomeDir)
		viper.AddConfigPath(join)

		fmt.Printf("UserHomeDir: %v, join:%v\n", homeDir, join)

		// 将当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath(".")

		// 设置配置文件格式为 yaml
		viper.SetConfigType("yaml")

		// 设置配置文件名称
		viper.SetConfigName(defaultConfigName)
	}

	// 读取匹配的环境变量
	viper.AutomaticEnv()

	// 读取环境变量的前缀为 MINIBLOG，如果是 miniblog，则自动转变为大写
	viper.SetEnvPrefix("MINIBLOG")

	// 将 viper.Get(key) key 字符串中的 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则再注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read viper configuration file", "err", err)
	}

	// 打印 viper 当前使用的配置文件，方便 Debug
	log.Debugw("Using config file", "file", viper.ConfigFileUsed())
}

// logOptions 从 viper 中读取日志配置，构建 `*log2.Options` 并返回
// ⚠️ `viper.Get<Type>()` 中 key 的名字需要使用 `.` 分割，以跟 YAML 中保持相同的锁进
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}
