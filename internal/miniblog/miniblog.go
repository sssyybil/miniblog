package miniblog

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"miniblog/internal/pkg/log"
)

var cfgFile string

// NewMiniBlogCommand 创建一个 *cobra.Command 对象，可通过 Command 对象的 Execute 方法来启动程序
func NewMiniBlogCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "miniblog",                          // 指定命令的名字，该名字会出现在帮助信息中
		Short: "A short and sweet blogging system", // 命令的简短描述
		Long: `A good Go practical project, used to create user with basic information.

Find more miniblog information at:
	https://github.com/marmotedu/miniblog#readme`, // 命令的详细描述
		SilenceUsage: true, // 命令出错时，不打印帮助信息。即可以保证命令出错时一眼就能看到错误信息
		RunE: func(cmd *cobra.Command, args []string) error { // 指定调用 cmd.Execute() 时，执行的 Run 函数，函数执行失败会返回错误信息
			// 初始化日志
			log.Init(logOptions())
			// Sync 将缓存中的日志刷新到磁盘中，以防日志丢失
			defer log.Sync()

			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error { // 设置命令运行时，不需要指定命令行参数
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	// 使 initConfig 函数在每个命令运行时都会被调用以读取配置
	cobra.OnInitialize(initConfig)

	// ⇩ 定义标志和配置设置

	// Cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the miniblog configuration file. Empty string for no configuration file.")

	// Cobra 也支持本地标志，本地标志只能在其所绑定的命令上使用
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return cmd
}

// run 函数是实际的业务代码入口函数
func run() error {
	fmt.Println("Hello MiniBlog~!!")

	// ↓ 测试 viper

	// 打印所有的配置项及值
	settings, _ := json.Marshal(viper.AllSettings())
	log.Infow(string(settings))

	// 打印 db 用户名
	log.Infow(viper.GetString("db.username"))

	return nil
}
