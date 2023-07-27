package version

import (
	"encoding/json"
	"fmt"
	"github.com/gosuri/uitable"
	"runtime"
)

/**
可分享给第三方开发者使用～
*/

var (
	// GitVersion 是语义化的版本号
	GitVersion = "v0.0.0-master+$Format:%h$"

	// BuildDate TODO 2023/7/25 14:13 sun: 为什么使用这种时间格式❓
	// BuildDate 是 ISO8601 格式的构建时间, $(date -u +'%Y-%m-%dT%H:%M:%SZ') 命令的输出
	BuildDate = "1970-01-01T00:00:00Z"

	// GitCommit 是 Git 的 SHA1 值，`$ git rev-parse HEAD` 命令的输出
	GitCommit = "$Format:%H$"

	// GitTreeState 代表构建时 Git 仓库的状态，可能的值有：clean、dirty
	GitTreeState = ""
)

// Info 包含了版本信息
type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String 返回人性化的版本信息字符串
func (info Info) String() string {
	if bytes, err := info.Text(); err == nil {
		return string(bytes)
	}
	return info.GitVersion
}

func (info Info) ToJSON() string {
	bytes, err := json.Marshal(info)
	if err != nil {
		fmt.Printf("ToJSON error: %s", err)
	}
	return string(bytes)
}

// Text 将版本信息编码为 UTF-8 格式的文本，并返回。
func (info Info) Text() ([]byte, error) {
	// github 第三方库 uitable（https://github.com/gosuri/uitable），可在终端应用程序中以表格形式表示数据。
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("gitVersion:", info.GitVersion) // addRow 函数内部做了加锁和解锁的操作
	table.AddRow("gitCommit:", info.GitCommit)
	table.AddRow("gitTreeState:", info.GitTreeState)
	table.AddRow("buildDate:", info.BuildDate)
	table.AddRow("goVersion:", info.GoVersion)
	table.AddRow("compiler:", info.Compiler)
	table.AddRow("platform:", info.Platform)

	return table.Bytes(), nil
}

// Get 返回详尽的代码库版本信息，用来标明二进制文件由哪个版本的代码构建
func Get() Info {
	return Info{
		GitVersion:   GitVersion,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler, // 常量值 “gc”
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
