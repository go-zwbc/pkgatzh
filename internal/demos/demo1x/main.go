// Package main demonstrates basic usage of pkgatzh
// Shows how to retrieve package path, name, import path, and module path information
//
// main 演示 pkgatzh 的基础用法
// 展示如何获取包路径、包名、导入路径和模块路径信息
package main

import (
	"fmt"

	"github.com/go-zwbc/pkgatzh"
)

func main() {
	// Demo: Retrieve package information (演示：获取包信息)

	// Create instance (创建实例)
	info := pkgatzh.NewT位置信息()

	// Print package information (打印包信息)
	fmt.Println("=== Package Information (包信息) ===")
	fmt.Println("Filesystem Path (文件系统路径):", info.P目录路径)
	fmt.Println("Package Name (包声明名):", info.N包的名称)
	fmt.Println("Import Path (引用路径):", info.I引用路径)
	fmt.Println("Module Path (项目模块):", info.M项目模块)
}
