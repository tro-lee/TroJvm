package classpath

import (
	"bytes"
	"errors"
	"strings"
)

// CompositeEntry 复合Entry
// 用于表示由更小的Entry组成的类路径
type CompositeEntry []Entry

// newCompositeEntry 构造函数
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := CompositeEntry{}
	// 遍历路径列表
	for _, path := range strings.Split(pathList, pathListSeparator) {
		// 调用newEntry函数创建Entry
		entry := newEntry(path)
		// 将创建的Entry添加到CompositeEntry中
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

// 实现Entry接口
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		// 依次调用每一个子路径的readClass方法，如果成功读取到class数据，返回
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil // 读取成功，返回
		}
	}
	return nil, nil, errors.New("类没找到: " + className)
}
func (self CompositeEntry) String() string {
	data := bytes.Buffer{}
	for i, entry := range self {
		// 依次调用每一个子路径的String方法，拼接结果
		if i > 0 {
			data.WriteString(pathListSeparator)
		}
		data.WriteString(entry.String())
	}
	return data.String()
}
