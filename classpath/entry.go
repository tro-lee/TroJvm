package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

// Entry 类路径项
// 举例作用：假设要读取java.lang.Object类，传入参数应该是java/lang/Object.class，
// 返回读取到的字节数据，最终定位到class文件的Entry
type Entry interface {
	// readClass 用来寻找和加载class文件
	readClass(className string) ([]byte, Entry, error)
	// String 类似返回变量的字符串表示
	String() string
}

// newEntry 创建新的Entry
func newEntry(path string) Entry {
	// 如果包含分隔符，说明是复合Entry
	// 例如: java -cp a.jar;b.jar;c.jar; com.example.Main
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	// 如果是*结尾，说明是通配符Entry
	// 例如: java -cp * com.example.Main
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	// 如果是zip或者jar包，说明是zip或者jar包Entry
	// 例如: java -cp a.jar com.example.Main
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	// 否则是目录Entry
	// 例如: java -cp com/example Main
	return newDirEntry(path)
}
