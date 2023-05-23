package main

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

// printUsage 打印命令行参数说明
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

// parseCmd 解析命令行参数
func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
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
