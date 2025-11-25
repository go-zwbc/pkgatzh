[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-zwbc/pkgatzh/release.yml?branch=main&label=BUILD)](https://github.com/go-zwbc/pkgatzh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-zwbc/pkgatzh)](https://pkg.go.dev/github.com/go-zwbc/pkgatzh)
[![Coverage Status](https://img.shields.io/coveralls/github/go-zwbc/pkgatzh/main.svg)](https://coveralls.io/github/go-zwbc/pkgatzh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/go-zwbc/pkgatzh.svg)](https://github.com/go-zwbc/pkgatzh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zwbc/pkgatzh)](https://goreportcard.com/report/github.com/go-zwbc/pkgatzh)

# pkgatzh

Chinese-named package with comprehensive package information extraction capabilities

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Main Features

📁 **Path Extraction**: Get absolute filesystem path to the invoking package
📦 **Name Extraction**: Extract package declaration name from source code
🌐 **Import Path**: Retrieve complete import path used in import statements
🏗️ **Module Path**: Extract module path defined in go.mod
🎯 **Context Awareness**: Uses runtime stack analysis to determine invocation context
🔧 **Chinese Naming**: Intuitive Chinese-named struct fields following "English prefix + 4 Chinese characters" pattern
✨ **Simple Integration**: Lightweight dependencies wrapping yyle88/runpath, yyle88/syntaxgo, and golang.org/x/tools/go/packages

## Installation

```bash
go get github.com/go-zwbc/pkgatzh
```

## Usage

### Basic Package Information Extraction

This example demonstrates extracting comprehensive package information.

```go
package main

import (
	"fmt"

	"github.com/go-zwbc/pkgatzh"
)

func main() {
	// Create instance
	info := pkgatzh.NewT位置信息()

	// Print package information
	fmt.Println("=== Package Information ===")
	fmt.Println("Filesystem Path:", info.P目录路径)
	fmt.Println("Package Name:", info.N包的名称)
	fmt.Println("Import Path:", info.I引用路径)
	fmt.Println("Module Path:", info.M项目模块)
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

## API Reference

### Type Definition

| Type | Description (EN) | 描述 (ZH) |
|------|------------------|-----------|
| `T位置信息` | Struct containing comprehensive package location information | 包含全面包位置信息的结构体 |

### Creation Function

| Function | Description (EN) | 描述 (ZH) |
|----------|------------------|-----------|
| `NewT位置信息()` | Creates new instance with comprehensive package information | 创建包含全面包信息的新实例 |

### Struct Fields

| Field | Description (EN) | 描述 (ZH) |
|-------|------------------|-----------|
| `P目录路径 string` | Absolute filesystem path to the package (e.g., "/path/to/pkg") | 包目录在文件系统中的绝对路径（如："/path/to/pkg"） |
| `N包的名称 string` | Package name from package declaration (e.g., "main"), not import path | package 声明中的包名（如："main"），不是导入路径 |
| `I引用路径 string` | Import path used in import statements (e.g., "github.com/go-zwbc/pkgatzh") | import 语句中使用的引用路径（如："github.com/go-zwbc/pkgatzh"） |
| `M项目模块 string` | Module path defined in go.mod (e.g., "github.com/go-zwbc/pkgatzh") | go.mod 中定义的模块路径（如："github.com/go-zwbc/pkgatzh"） |

## Examples

### **Testing Context**

```go
func TestSomething(t *testing.T) {
    info := pkgatzh.NewT位置信息()
    testDataPath := filepath.Join(info.P目录路径, "testdata")
    // Load test data from testDataPath
}
```

### **Configuration Paths**

```go
func init() {
    info := pkgatzh.NewT位置信息()
    configPath := filepath.Join(info.P目录路径, "config.yaml")
    // Load configuration from package-based path
}
```

### **Package Identification**

```go
func NewService() *Service {
    info := pkgatzh.NewT位置信息()
    return &Service{
        name: info.N包的名称,
        path: info.P目录路径,
        importPath: info.I引用路径,
        modulePath: info.M项目模块,
    }
}
```

## Implementation Details

### Filesystem Path Extraction
- Uses `yyle88/runpath` to get runtime path information
- `runpath.PARENT.Skip(1)` gets the parent path of the invoking package
- Returns absolute path to the package on filesystem

### Package Name Extraction
- Uses `yyle88/syntaxgo` to parse Go source code
- `syntaxgo.GetPkgName(runpath.Skip(1))` extracts package name from source
- Returns the package name as declared in Go source

### Import Path and Module Path Extraction
- Uses `golang.org/x/tools/go/packages` to load package metadata
- Supports both production packages and test packages through intelligent matching
- Extracts complete import path and module path information

### Context Detection
- Runtime stack analysis determines the invocation point
- `Skip(1)` moves up one stack frame to get the context info
- Handles different package nesting situations without issues

## Naming Conventions

- `T` prefix: Type definitions (T位置信息)
- `New` prefix: Creation functions (NewT位置信息)
- `P` prefix: Path fields (P目录路径)
- `N` prefix: Name fields (N包的名称)
- `I` prefix: Import path fields (I引用路径)
- `M` prefix: Module path fields (M项目模块)
- Field naming: "English prefix + 4 Chinese characters" pattern (e.g., P目录路径, N包的名称)

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-20 04:26:32.402216 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->
