package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
	"time"
)

var (
	mu  sync.Mutex
	std = NewLogger(NewOptions()) // 定义默认的全局 Logger
)

type ZapLogger struct {
	z *zap.Logger
}

// 🌻确保 zapLogger 实现了 Logger 接口，以下变量赋值，可以使错误在编译器被发现。该编程技巧在 Go 项目开发中被大量使用
var _ Logger = &ZapLogger{}

// Init 使用指定的选项初始化 Logger
func Init(opts *Options) {
	// TODO 2023/7/24 14:41 sun: 此处加锁的原因？
	mu.Lock()
	defer mu.Unlock()

	std = NewLogger(opts)
}

// NewLogger 根据传入的 opts 创建 Logger
func NewLogger(opts *Options) *ZapLogger {
	if opts == nil {
		opts = NewOptions()
	}

	// 创建一个默认的 encoder 配置。EncoderConfig 结构定义了日志信息在写入输出端之前的编码和格式化方式
	encoderConfig := zap.NewProductionEncoderConfig()
	// 默认值为 “msg”，自定义 MessageKey 为 message，message 语义更明确
	encoderConfig.MessageKey = "message"
	// 默认值为 “ts”，自定义 TimeKey 为 timestamp，timestamp 语义更明确
	encoderConfig.TimeKey = "timestamp"
	// 指定时间序列化函数，将时间序列化为 `2006-01-02 15:04:05.000` 格式，更易读
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Format("2006-01-02 15:04:05.000"))
	}
	// 指定 time.Duration 序列化函数，将 time.Duration 序列化为经过的毫秒数的浮点数（默认为秒），毫秒 数比默认的秒数更精确
	encoderConfig.EncodeDuration = func(duration time.Duration, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendFloat64(float64(duration) / float64(time.Millisecond))
	}

	// 将文本的日志级别，例如 info 转换为 zapcore.Level 类型
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		// 若指定了非法的日志级别，则默认使用 info 级别
		zapLevel = zapcore.InfoLevel
	}

	cfg := &zap.Config{
		DisableCaller:     opts.DisableCaller,             // 是否在日志中显示调用日志所在的文件和行号，例如：`"caller":"miniblog/miniblog.go:75"`
		DisableStacktrace: opts.DisableStacktrace,         // 是否禁止 panic 及以上级别打印堆栈信息
		Level:             zap.NewAtomicLevelAt(zapLevel), // 指定日志级别
		Encoding:          opts.Format,                    // 指定日志显示格式，可选值：console, json
		EncoderConfig:     encoderConfig,
		OutputPaths:       opts.OutputPaths,   // 指定日志输出位置
		ErrorOutputPaths:  []string{"stderr"}, // 设置 zap 内部错误输出位置
	}

	// TODO 2023/7/24 20:01 sun: 使用 cfg 创建 *zap.Logger 对象。参数含义含义❓因为是自定义封装的 zap 包，所以在调用栈中跳过的调用深度要加 1
	z, err := cfg.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		log.Fatalln(err)
	}

	logger := &ZapLogger{z: z}

	return logger
}

// Logger 定义了 MiniBlog 项目的日志接口，该接口只包含了支持的日志记录方法。接口中的函数名采用了 zap 中的函数名
type Logger interface {
	Debugw(msg string, keyAndValues ...any)
	Infow(msg string, keyAndValues ...any)
	Warnw(msg string, keyAndValues ...any)
	Errorw(msg string, keyAndValues ...any)
	Panicw(msg string, keyAndValues ...any)
	Fatalw(msg string, keyAndValues ...any)
	Sync()
}

// Sync 调用底层 zap.Logger 的 Sync 方法，将缓存中的日志刷新到磁盘文件中，主程序需要在推出前调用 Sync
func Sync() {
	err := std.z.Sync()
	if err != nil {
		log.Printf("Sync function error: %v\n", err)
	}
}

// Debugw 输出 debug 级别的日志
func Debugw(msg string, keyAndValues ...any) {
	std.z.Sugar().Debugw(msg, keyAndValues)
}

// Infow 输出 info 级别的日志
func Infow(msg string, keyAndValues ...any) {
	std.z.Sugar().Infow(msg, keyAndValues)
}

// Warnw 输出 warn 级别的日志
func Warnw(msg string, keyAndValues ...any) {
	std.z.Sugar().Warnw(msg, keyAndValues)
}

// Errorw 输出 error 级别的日志
func Errorw(msg string, keyAndValues ...any) {
	std.z.Sugar().Errorw(msg, keyAndValues)
}

// Panicw 输出 panic 级别的日志
func Panicw(msg string, keyAndValues ...any) {
	std.z.Sugar().Panicw(msg, keyAndValues)
}

// Fatalw 输出 fatal 级别的日志
func Fatalw(msg string, keyAndValues ...any) {
	std.z.Sugar().Fatalw(msg, keyAndValues)
}

func (zl *ZapLogger) Debugw(msg string, keyAndValues ...any) {
	zl.z.Sugar().Debugw(msg, keyAndValues)
}

func (zl *ZapLogger) Infow(msg string, keyAndValues ...any) {
	zl.z.Sugar().Infow(msg, keyAndValues)
}

func (zl *ZapLogger) Warnw(msg string, keyAndValues ...any) {
	zl.z.Sugar().Warnw(msg, keyAndValues)
}

func (zl *ZapLogger) Errorw(msg string, keyAndValues ...any) {
	zl.z.Sugar().Errorw(msg, keyAndValues)
}

func (zl *ZapLogger) Panicw(msg string, keyAndValues ...any) {
	zl.z.Sugar().Panicw(msg, keyAndValues)
}

func (zl *ZapLogger) Fatalw(msg string, keyAndValues ...any) {
	zl.z.Sugar().Fatalw(msg, keyAndValues)
}

func (zl *ZapLogger) Sync() {
	err := zl.z.Sync()
	if err != nil {
		log.Printf("Sync function error: %v\n", err)
	}
}
