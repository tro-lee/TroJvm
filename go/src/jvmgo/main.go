package main

import (
	"TroJvm/go/src/jvmgo/classpath"
	"fmt"
	"strings"
)

func main() {
	//解析命令行
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	// 解析类路径
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	println("classpath:%s class:%s args:%v\n", cp, cmd.class, cmd.args)

	// 把.全部替换成/
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}

	fmt.Printf("%v", classData)
}
