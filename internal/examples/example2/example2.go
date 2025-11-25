// Package example2 demonstrates more usage of pkgatzh package
// Shows how distinct packages can retrieve path, name, and import path data
//
// example2 包演示 pkgatzh 包的另一种用法
// 展示不同包如何获取路径、名称和导入路径数据
package example2

import "github.com/go-zwbc/pkgatzh"

// V位置信息 holds the path, name, and import path information of this package
// V位置信息 保存此包的路径、名称和导入路径信息
var V位置信息 = pkgatzh.NewT位置信息()
