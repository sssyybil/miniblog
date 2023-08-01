package miniblog

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"miniblog/internal/pkg/log"
	"miniblog/internal/pkg/middleware"
	"miniblog/pkg/version/verflag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

			// 如果 `--version=true`，则打印版本并退出
			verflag.PrintAdnExitIfRequested()

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

	// 添加 --version 版本信息
	verflag.AddFlags(cmd.PersistentFlags())

	return cmd
}

// run 函数是实际的业务代码入口函数
func run() error {

	// 设置 Gin 模式
	gin.SetMode(viper.GetString("runmode"))

	// 创建 Gin 引擎
	g := gin.New()

	// gin.Recover 中间件，用来捕获任何 panic 并恢复
	middlewares := []gin.HandlerFunc{gin.Recovery(), middleware.NoCache, middleware.Cors, middleware.Secure, middleware.RequestID()}

	g.Use(middlewares...)

	// 注册 404 Handler
	g.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"code": 10003, "message": "Page not found."}) // 将结果序列化为 JSON 格式放入 ResponseBody 中
	})

	// 注册 /health Handler
	g.GET("/health", func(context *gin.Context) {
		log.C(context).Infow("Health function called")

		context.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// TODO 2023/7/27 15:08 sun: 为什么不直接使用 g.Run()
	// 创建 HTTP Server 实例
	server := &http.Server{Addr: viper.GetString("addr"), Handler: g}

	logger, _ := zap.NewProduction()

	logger.Sugar().Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))
	log.Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))

	go func() {
		// 调用 server.shutdown() 方法时，Server、ListenAndServe、ListenAndServeTLS 方法会立刻返回 ErrServerClosed 错误，该错误为服务器关闭时的正常报错行为
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

	// 等待中断信号，优雅的关闭服务器（10s 超时）
	quit := make(chan os.Signal, 1)
	/**
	此处不阻塞
	kill 默认会发送 SIGINT 信号
	kill -2 发送 SIGTERM 信号（或 Ctrl+C）
	kill -9 会发送 SIGKILL 信号，但无法被捕获，所以不添加在此处
	*/
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞在此，当接收到以上两种信号中的某一个时才会继续往下面进行
	<-quit
	log.Infow("Shutting down server...")

	// 创建 ctx 用于通知服务器 goroutine，它有 10 秒时间完成当前正在处理的请求
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	if err := server.Shutdown(ctx); err != nil {
		log.Errorw("Insecure Server forced to shutdown", "err", err)
		return err
	}

	log.Infow("Server existing")

	return nil
}
