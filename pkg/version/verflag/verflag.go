package verflag

import (
	"fmt"
	"github.com/spf13/pflag"
	"miniblog/pkg/version"
	"os"
	"strconv"
)

type VersionValue int

const (
	VersionFalse VersionValue = 0
	VersionTrue  VersionValue = 1
	VersionRaw   VersionValue = 2
)

const (
	strRawVersion   = "raw"
	versionFlagName = "version"
)

var versionFlag = Version(versionFlagName, VersionFalse, "Print version information and quit.")

// TODO 2023/7/25 15:13 sun: 下述方法中，versionValue 为什么要设置为指针类型呢❓

func (v *VersionValue) IsBoolFlag() bool {
	return true
}

func (v *VersionValue) Get() any {
	return v
}

// String 实现了 pflag.Value 接口中的 String 方法
func (v *VersionValue) String() string {
	if *v == VersionRaw {
		return strRawVersion
	}
	return fmt.Sprintf("%v", *v == VersionTrue)
}

// Set 实现了 pflag.Value 接口中的 Set 方法
func (v *VersionValue) Set(s string) error {
	// TODO 2023/7/25 15:41 sun: 解释❓
	if s == strRawVersion {
		*v = VersionRaw
		return nil
	}

	boolVal, err := strconv.ParseBool(s)
	if boolVal {
		*v = VersionTrue
	} else {
		*v = VersionFalse
	}
	return err
}

// Type 实现了 pflag.Value 接口中的 Type 方法
func (v *VersionValue) Type() string {
	return "version"
}

// VersionVar 定义了一个具有指定名称和用法的标志
func VersionVar(p *VersionValue, name string, value VersionValue, usage string) {
	*p = value // p == 0
	pflag.Var(p, name, usage)
	// `--version` 等价于 `--version=true`
	pflag.Lookup(name).NoOptDefVal = "true"
}

// Version 包装了 VersionVar 函数. ("version",0,"Print version information and quit.")
func Version(name string, value VersionValue, usage string) *VersionValue {
	p := new(VersionValue)
	// TODO 2023/7/25 15:47 sun: p == nil ❓
	VersionVar(p, name, value, usage)
	return p
}

// AddFlags 在任意 FlagSet 上注册这个包的标志，这样它们指向与全局标志相同的值
func AddFlags(fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(versionFlagName))
}

// PrintAdnExitIfRequested 检查是否传递了`--version` 标志，如果是，则打印版本信息并退出
func PrintAdnExitIfRequested() {

	// 根据 *versionFlag 值不同，输出不同格式的版本信息
	if *versionFlag == VersionRaw {
		fmt.Printf("%#v\n", version.Get())
		os.Exit(0)
	} else if *versionFlag == VersionTrue {
		fmt.Printf("%s\n", version.Get())
		os.Exit(0)
	}

	// TODO 2023/7/25 16:01 sun: 改成 switch ❓
}
