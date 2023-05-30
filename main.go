package main

import (
	"TroJvm/classpath"
	"TroJvm/interpreter"
	"TroJvm/rtda/heap"
	"fmt"
	"strings"
)

func main() {
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
	// 解析获得ClassPath
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	// 创建类加载器
	loader := heap.NewClassLoader(cp)

	// 先加载主类和主方法，存在内存中，通过方法获取类
	className := strings.Replace(cmd.class, ".", "/", -1)
	// 加载类获得数据
	mainClass := loader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()

	if mainMethod != nil {
		// 执行主函数
		interpreter.Interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
