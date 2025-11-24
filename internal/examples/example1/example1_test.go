package example1_test

import (
	"testing"

	"github.com/go-zwbc/pkgatzh/internal/examples/example1"
	"github.com/yyle88/neatjson/neatjsons"
)

// TestExample tests the example1 package path and name information extraction
// Verifies that V位置信息 contains valid package data
//
// TestExample 测试 example1 包的路径和名称信息提取
// 验证 V位置信息 包含有效的包数据
func TestExample(t *testing.T) {
	// Log the package information in JSON format (以 JSON 格式记录包信息)
	t.Log(neatjsons.S(example1.V位置信息))
}
