package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

// ZipEntry zip包Entry
// 用于表示zip或者jar格式的类路径
type ZipEntry struct {
	absPath string
}

// newZipEntry 构造函数
func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path) // 获取绝对路径
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

// 实现Entry接口
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath) // 打开zip包
	if err != nil {
		return nil, nil, err
	}
	defer r.Close() // 关闭zip包

	// 遍历zip包中的文件
	for _, f := range r.File {
		// 如果找到了class文件，读取并返回
		if f.Name == className {
			// 打开文件，创建一个ReadCloser
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()

			// 读取文件
			data, err := io.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}

			return data, self, err
		}
	}
	return nil, nil, errors.New("类没找到: " + className)
}
func (self *ZipEntry) String() string {
	return self.absPath // 直接返回地址
}
