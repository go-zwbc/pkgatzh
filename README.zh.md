[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/go-zwbc/pkgatzh/release.yml?branch=main&label=BUILD)](https://github.com/go-zwbc/pkgatzh/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-zwbc/pkgatzh)](https://pkg.go.dev/github.com/go-zwbc/pkgatzh)
[![Coverage Status](https://img.shields.io/coveralls/github/go-zwbc/pkgatzh/main.svg)](https://coveralls.io/github/go-zwbc/pkgatzh?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/go-zwbc/pkgatzh.svg)](https://github.com/go-zwbc/pkgatzh/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zwbc/pkgatzh)](https://goreportcard.com/report/github.com/go-zwbc/pkgatzh)

# pkgatzh

使用中文命名提供全面的包信息提取能力

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 主要特性

📁 **路径获取**: 获取调用包的文件系统绝对路径
📦 **名称提取**: 从源代码提取 package 声明名称
🌐 **引用路径**: 获取 import 语句中使用的完整引用路径
🏗️ **模块路径**: 提取 go.mod 中定义的模块路径
🎯 **上下文感知**: 使用运行时栈分析确定调用上下文
🔧 **中文命名**: 直观的中文命名结构体字段，遵循"字母 + 4 汉字"模式
✨ **简单集成**: 轻量依赖封装 yyle88/runpath、yyle88/syntaxgo 和 golang.org/x/tools/go/packages

## 注意事项

⚠️ 这个包依赖源码环境（需要 .go 文件和 go.mod 存在），仅适用于开发和测试场景。编译后的二进制文件在源码目录不存在时无法正常工作。

## 安装

```bash
go get github.com/go-zwbc/pkgatzh
```

## 使用方法

### 基础包信息获取

此示例演示获取全面的包信息。

```go
package main

import (
	"fmt"

	"github.com/go-zwbc/pkgatzh"
)

func main() {
	// 创建实例
	info := pkgatzh.NewT位置信息()

	// 显示包信息
	fmt.Println("=== 包信息 ===")
	fmt.Println("文件系统路径:", info.P目录路径)
	fmt.Println("包名:", info.N包的名称)
	fmt.Println("引用路径:", info.I引用路径)
	fmt.Println("模块路径:", info.M项目模块)
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

## API 参考

### 类型定义

| 类型 | 描述 (ZH) | Description (EN) |
|------|-----------|------------------|
| `T位置信息` | 包含全面包位置信息的结构体 | Struct containing comprehensive package location information |

### 创建函数

| 函数 | 描述 (ZH) | Description (EN) |
|------|-----------|------------------|
| `NewT位置信息()` | 创建包含全面包信息的新实例 | Creates new instance with comprehensive package information |

### 结构体字段

| 字段 | 描述 (ZH) | Description (EN) |
|------|-----------|------------------|
| `P目录路径 string` | 包目录在文件系统中的绝对路径（如："/path/to/pkg"） | Absolute filesystem path to the package (e.g., "/path/to/pkg") |
| `N包的名称 string` | package 声明中的包名（如："main"），不是导入路径 | Package name from package declaration (e.g., "main"), not import path |
| `I引用路径 string` | import 语句中使用的引用路径（如："github.com/go-zwbc/pkgatzh"） | Import path used in import statements (e.g., "github.com/go-zwbc/pkgatzh") |
| `M项目模块 string` | go.mod 中定义的模块路径（如："github.com/go-zwbc/pkgatzh"） | Module path defined in go.mod (e.g., "github.com/go-zwbc/pkgatzh") |

## 示例

### **测试上下文**

```go
func TestSomething(t *testing.T) {
    info := pkgatzh.NewT位置信息()
    testDataPath := filepath.Join(info.P目录路径, "testdata")
    // 从 testDataPath 加载测试数据
}
```

### **配置路径**

```go
func init() {
    info := pkgatzh.NewT位置信息()
    configPath := filepath.Join(info.P目录路径, "config.yaml")
    // 从基于包的路径加载配置
}
```

### **包标识**

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

## 实现细节

### 文件系统路径获取
- 使用 `yyle88/runpath` 获取运行时路径信息
- `runpath.PARENT.Skip(1)` 获取调用包的父级路径
- 返回包在文件系统上的绝对路径

### 包名提取
- 使用 `yyle88/syntaxgo` 解析 Go 源代码
- `syntaxgo.GetPkgName(runpath.Skip(1))` 从源码提取包名
- 返回在 Go 源码中声明的包名

### 引用路径和模块路径提取
- 使用 `golang.org/x/tools/go/packages` 加载包元数据
- 通过智能匹配支持生产包和测试包
- 提取完整的引用路径和模块路径信息

### 上下文检测
- 运行时栈分析确定调用点
- `Skip(1)` 向上移动一个栈帧以获取上下文信息
- 在不同包嵌套情况下都能正常处理

## 命名规范

- `T` 前缀: 类型定义 (T位置信息)
- `New` 前缀: 创建函数 (NewT位置信息)
- `P` 前缀: 路径字段 (P目录路径)
- `N` 前缀: 名称字段 (N包的名称)
- `I` 前缀: 引用路径字段 (I引用路径)
- `M` 前缀: 模块路径字段 (M项目模块)
- 字段命名: "字母 + 4 汉字"模式（如：P目录路径、N包的名称）

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-20 04:26:32.402216 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->
