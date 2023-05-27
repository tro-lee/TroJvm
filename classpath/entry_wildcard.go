package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// WildcardEntry 通配符Entry
// 由于wildcardEntry也是CompositeEntry，故不再定义新类型
func newWildcardEntry(path string) CompositeEntry {
	// 去掉路径末端*
	baseDir := path[:len(path)-1]
	// 生成CompositeEntry
	compositeEntry := CompositeEntry{}

	// 遍历baseDir
	// 继承接口，使用walkFn
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // 如果出错，返回
		}
		// info.IsDir()判断是否是目录, path != baseDir判断是否是baseDir本身
		if info.IsDir() && path != baseDir {
			// 跳过子目录
			return filepath.SkipDir
		}
		// 如果是jar文件，创建ZipEntry
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	// 使用相对路径
	filepath.Walk(baseDir, walkFn)

	return compositeEntry
}
