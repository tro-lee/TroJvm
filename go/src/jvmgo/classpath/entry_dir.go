package classpath

import (
	"os"
	"path/filepath"
)

// DirEntry 目录Entry
// 用于表示目录形式的类路径
type DirEntry struct {
	absDir string // 目录的绝对路径
}

// newDirEntry 构造函数
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path) // 获取目录的绝对路径
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

// 实现Entry接口
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className) // 拼上地址
	data, err := os.ReadFile(fileName)                // 读取文件
	return data, self, err

}
func (self *DirEntry) String() string {
	return self.absDir //直接返回目录
}
