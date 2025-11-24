// Package main demonstrates basic usage of pkgatzh
// Shows how to retrieve package path and name information
//
// main 演示 pkgatzh 的基础用法
// 展示如何获取包路径和包名信息
package main

import (
	"fmt"

	"github.com/go-zwbc/pkgatzh"
)

func main() {
	// Demo: Retrieve package path and name (演示：获取包路径和包名)

	// Create instance (创建实例)
	info := pkgatzh.NewT位置信息()

	// Print package information (打印包信息)
	fmt.Println("=== Package Information (包信息) ===")
	fmt.Println("Package Path (包路径):", info.P路径)
	fmt.Println("Package Name (包名):", info.N包名)
}
