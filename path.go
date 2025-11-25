// Package pkgatzh provides Chinese-named APIs to retrieve comprehensive package information
// Retrieves filesystem path, package name, import path, and module path information
// Wraps yyle88/runpath and yyle88/syntaxgo with golang.org/x/tools/go/packages
// Designed with intuitive Chinese field names following "letter + 4 characters" pattern
//
// pkgatzh 包提供中文命名的 API 来获取全面的包信息
// 获取文件系统路径、包名、引用路径和项目模块路径信息
// 封装 yyle88/runpath、yyle88/syntaxgo 和 golang.org/x/tools/go/packages
// 采用直观的中文字段名，遵循"字母 + 4 汉字"命名模式
package pkgatzh

import (
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"github.com/yyle88/syntaxgo"
	"golang.org/x/tools/go/packages"
)

// T位置信息 represents comprehensive package location information
// Contains filesystem path, package declaration name, import path, and module path
// Each field follows "letter + 4 Chinese characters" naming convention
//
// T位置信息 表示全面的包位置信息
// 包含文件系统路径、包声明名称、引用路径和项目模块路径
// 每个字段遵循"字母 + 4 汉字"命名规范
type T位置信息 struct {
	P目录路径 string // P目录路径 is the absolute filesystem path to the package (P目录路径 是包目录在文件系统中的绝对路径)
	N包的名称 string // N包的名称 is the package name from package declaration, not the import path (N包的名称 是 package 声明中的包名，不是引用路径)
	I引用路径 string // I引用路径 is the import path used in import statements (I引用路径 是在 import 语句中使用的引用路径)
	M项目模块 string // M项目模块 is the module path defined in go.mod (M项目模块 是 go.mod 中定义的模块路径)
}

// NewT位置信息 creates new T位置信息 instance through retrieving invoking package's information
// Uses runtime stack analysis to determine invoking package's location at execution time
// Leverages golang.org/x/tools/go/packages to extract complete package metadata
// Handles both production packages and test packages through intelligent package matching
//
// Field Information:
// - P目录路径: Absolute filesystem path (e.g., "/path/to/mypackage")
// - N包的名称: Package name from declaration (e.g., "main", "mypackage"), not import path
// - I引用路径: Import path used in statements (e.g., "github.com/go-zwbc/pkgatzh")
// - M项目模块: Module path defined in go.mod (e.g., "github.com/go-zwbc/pkgatzh")
//
// NewT位置信息 通过获取调用者的包信息来创建新的 T位置信息 实例
// 使用运行时栈分析来动态确定调用包的位置
// 利用 golang.org/x/tools/go/packages 提取完整的包元数据
// 通过智能包匹配处理生产包和测试包
//
// 字段信息：
// - P目录路径：文件系统的绝对路径（例如："/path/to/mypackage"）
// - N包的名称：package 声明中的包名（例如："main"、"mypackage"），不是引用路径
// - I引用路径：import 语句中使用的引用路径（例如："github.com/go-zwbc/pkgatzh"）
// - M项目模块：go.mod 中定义的模块路径（例如："github.com/go-zwbc/pkgatzh"）
func NewT位置信息() *T位置信息 {
	absPath := runpath.PARENT.Skip(1)                   // Get invoking package's absolute path on filesystem (获取调用者包的绝对目录路径)
	pkgName := syntaxgo.GetPkgName(runpath.Skip(1))     // Get package name from source code (从源文件获取实际包名)
	pkg := must.Full(getPackage(absPath, pkgName))      // Must get valid package matching the name (必须获取匹配名称的有效包)
	importPath := must.Nice(pkg.PkgPath)                // Must get import path (必须获取引用路径)
	modulePath := must.Nice(must.Full(pkg.Module).Path) // Must get module path (必须获取项目模块路径)
	return &T位置信息{
		P目录路径: absPath,                      // Absolute filesystem path to invoking package (调用者包目录的文件系统绝对路径)
		N包的名称: must.Sane(pkg.Name, pkgName), // Package name from "package xxx" declaration (从 "package xxx" 声明中提取的包名)
		I引用路径: importPath,                   // Import path from go.mod and package metadata (从 go.mod 和包元数据获取的引用路径)
		M项目模块: modulePath,                   // Module path from go.mod (从 go.mod 获取的项目模块路径)
	}
}

// getPackage retrieves package information from filesystem path matching the package name
// Loads packages using golang.org/x/tools/go/packages with test package support
// Searches through loaded packages to find the one with matching package name
// Returns matched package, using first package as fallback when no exact match exists
//
// getPackage 从文件系统路径获取匹配包名的包信息
// 使用 golang.org/x/tools/go/packages 加载包，支持测试包
// 在加载的包中搜索匹配包名的那个
// 返回匹配的包，如果没有精确匹配则返回第一个包作为降级
func getPackage(pkgPath string, pkgName string) *packages.Package {
	cfg := &packages.Config{
		Mode:  packages.NeedName | packages.NeedModule, // Request package name and module info (请求包名和模块信息)
		Dir:   pkgPath,                                 // Set working path (设置工作路径)
		Tests: true,                                    // Include test packages to find both production and test packages (包含测试包以查找普通包和测试包)
	}
	pkgs := rese.A1(packages.Load(cfg, ".")) // Load packages from current path, discard errors (从当前路径加载包，丢弃错误)
	must.Have(pkgs)                          // Assert packages loaded without issues (断言包加载成功)

	// Find the package matching the given name (查找匹配给定名称的包)
	for _, pkg := range pkgs {
		if pkg.Name == pkgName {
			return pkg
		}
	}

	// If no match found, return first package as fallback (如果没找到匹配的，返回第一个包作为降级)
	return must.Full(pkgs[0])
}
