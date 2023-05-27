package classpath

import (
	"os"
	"path/filepath"
)

// ClassPath 结构体，存放三种类路径
type ClassPath struct {
	// 启动类路径
	bootClasspath Entry
	// 扩展类路径
	extClasspath Entry
	// 用户类路径
	userClasspath Entry
}

// Parse 解析类路径
func Parse(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	// 解析启动类路径和扩展类路径
	cp.parseBootAndExtClasspath(jreOption)
	// 解析用户类路径
	cp.parseUserClasspath(cpOption)

	return cp
}

// parseBootAndExtClasspath 解析启动类路径和扩展类路径
func (self *ClassPath) parseBootAndExtClasspath(jreOption string) {
	// 对jre目录进行判断
	jreDir := getJreDir(jreOption)

	// jre/lib/* 解析启动类路径
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/* 解析扩展类路径
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

// 辅助函数 getJreDir 获取jre目录
// 优先使用用户输入-Xjre选项作为jre目录，
// 如果没有，就用当前目录下的jre目录，如果找不到，就尝试使用JAVA_HOME环境变量
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("找不到jre目录！")
}

// 辅助函数 exists 判断文件是否存在
func exists(path string) bool {
	// 如果返回的错误为nil，说明文件或者目录存在
	if _, err := os.Stat(path); err != nil {
		// 再次确认错误类型是否为不存在，还是其他
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// parseUserClasspath 解析用户类路径
func (self *ClassPath) parseUserClasspath(cpOption string) {
	// 如果用户没有提供-classpath/-cp选项，则使用当前目录作为用户类路径
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

// ReadClass 读取
func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	// 将类名转换为相对路径
	className = className + ".class"
	// 依次从启动类路径、扩展类路径和用户类路径中搜索class文件
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

// String 转化为字符串
func (self *ClassPath) String() string {
	return self.userClasspath.String()
}
