[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-zwbc/pkgatzh/release.yml?branch=main&label=BUILD)](https://github.com/go-zwbc/pkgatzh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-zwbc/pkgatzh)](https://pkg.go.dev/github.com/go-zwbc/pkgatzh)
[![Coverage Status](https://img.shields.io/coveralls/github/go-zwbc/pkgatzh/main.svg)](https://coveralls.io/github/go-zwbc/pkgatzh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/go-zwbc/pkgatzh.svg)](https://github.com/go-zwbc/pkgatzh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zwbc/pkgatzh)](https://goreportcard.com/report/github.com/go-zwbc/pkgatzh)

# pkgatzh

Chinese-named package for retrieving package path and name information

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Main Features

📁 **Path Extraction**: Get absolute path to the package that invokes the function
📦 **Name Extraction**: Extract package name from source code
🎯 **Context Awareness**: Uses runtime stack information to determine context
🔧 **Chinese Naming**: Intuitive Chinese-named struct and method APIs
✨ **Simple Integration**: Lightweight dependencies wrapping yyle88/runpath and yyle88/syntaxgo

## Installation

```bash
go get github.com/go-zwbc/pkgatzh
```

## Usage

### Basic Package Information Extraction

This example demonstrates extracting package path and name information.

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
	fmt.Println("Package Path:", info.P路径)
	fmt.Println("Package Name:", info.N包名)
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

## API Reference

### Type Definition

| Type | Description (EN) | 描述 (ZH) |
|------|------------------|-----------|
| `T位置信息` | Struct containing package path and name | 包含包路径和包名的结构体 |

### Creation Function

| Function | Description (EN) | 描述 (ZH) |
|----------|------------------|-----------|
| `NewT位置信息()` | Creates new instance with package path and name | 创建包含包路径和包名的新实例 |

### Struct Fields

| Field | Description (EN) | 描述 (ZH) |
|-------|------------------|-----------|
| `P路径 string` | Absolute filesystem path to the package DIR (e.g., "/path/to/pkg") | 包目录在文件系统中的绝对路径（如："/path/to/pkg"） |
| `N包名 string` | Package name from package declaration (e.g., "main"), not import path | package 声明中的包名（如："main"），不是导入路径 |

## Examples

### **Testing Context**

```go
func TestSomething(t *testing.T) {
    info := pkgatzh.NewT位置信息()
    testDataDIR := filepath.Join(info.P路径, "testdata")
    // Load test data from testDataDIR
}
```

### **Configuration Paths**

```go
func init() {
    info := pkgatzh.NewT位置信息()
    configPath := filepath.Join(info.P路径, "config.yaml")
    // Load configuration from package-based path
}
```

### **Package Identification**

```go
func NewService() *Service {
    info := pkgatzh.NewT位置信息()
    return &Service{
        name: info.N包名,
        path: info.P路径,
    }
}
```

## Implementation Details

### Path Extraction
- Uses `yyle88/runpath` to get runtime path information
- `runpath.PARENT.Skip(1)` gets the parent DIR of the invoking package
- Returns absolute path to the package

### Name Extraction
- Uses `yyle88/syntaxgo` to parse Go source code
- `syntaxgo.GetPkgName(runpath.Skip(1))` extracts package name from source
- Returns the package name as declared in Go source

### Context Detection
- Runtime stack analysis determines the invocation point
- `Skip(1)` moves up one stack frame to get the context info
- Handles different package nesting situations without issues

## Naming Conventions

- `T` prefix: Type definitions (T位置信息)
- `New` prefix: Creation functions (NewT位置信息)
- `P` prefix: Path fields (P路径)
- `N` prefix: Name fields (N包名)

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
