package cmd

import (
	"flag"
	"fmt"
	"os"
)

// Cmd 定义java命令行参数
type Cmd struct {
	helpFlag    bool     //帮助信息
	versionFlag bool     //版本号
	cpOption    string   //指定用户类路径
	XjreOption  string   //指定jre目录
	class       string   //主类名
	args        []string //参数
}

func (c Cmd) HelpFlag() bool {
	return c.helpFlag
}

func (c Cmd) VersionFlag() bool {
	return c.versionFlag
}

func (c Cmd) CpOption() string {
	return c.cpOption
}

func (c Cmd) Class() string {
	return c.class
}

func (c Cmd) Args() []string {
	return c.args
}

// PrintUsage 打印命令行参数说明
func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

// ParseCmd 解析命令行参数
func ParseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = PrintUsage
	//设置命令行解析选项
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")

	//设置命令行参数
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}
