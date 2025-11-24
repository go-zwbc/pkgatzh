// Package pkgatzh provides Chinese-named functions to retrieve package path and name information.
// It wraps yyle88/runpath and yyle88/syntaxgo to provide intuitive Chinese-named APIs.
//
// pkgatzh 包提供中文命名的函数来获取包路径和包名信息。
// 它封装了 yyle88/runpath 和 yyle88/syntaxgo 以提供直观的中文命名 API。
package pkgatzh

import (
	"github.com/yyle88/runpath"
	"github.com/yyle88/syntaxgo"
)

// T位置信息 represents a package with its filesystem path and package name
// T位置信息 表示一个包及其文件系统路径和包声明名称
type T位置信息 struct {
	P路径 string // P路径 is the absolute filesystem path to the package DIR (P路径 是包目录在文件系统中的绝对路径)
	N包名 string // N包名 is the package name from package declaration, not the import path (N包名 是 package 声明中的包名，不是导入路径)
}

// NewT位置信息 creates a new T位置信息 instance through retrieving the invoking package's information.
// It uses runtime stack information to determine the invoking package's location.
// The P路径 field contains the absolute filesystem path (e.g., "/path/to/mypackage").
// The N包名 field contains the package name from package declaration (e.g., "main", "mypackage"), not the import path.
//
// NewT位置信息 通过获取调用者的包信息来创建新的 T位置信息 实例。
// 它使用运行时栈信息来确定调用包的位置。
// P路径 字段包含文件系统的绝对路径（例如："/path/to/mypackage"）。
// N包名 字段包含 package 声明中的包名（例如："main"、"mypackage"），不是导入路径。
func NewT位置信息() *T位置信息 {
	return &T位置信息{
		P路径: runpath.PARENT.Skip(1),               // Absolute filesystem path to invoking package's DIR (调用者包目录的文件系统绝对路径)
		N包名: syntaxgo.GetPkgName(runpath.Skip(1)), // Package name from "package xxx" declaration (从 "package xxx" 声明中提取的包名)
	}
}
