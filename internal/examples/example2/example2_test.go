package example2_test

import (
	"testing"

	"github.com/go-zwbc/pkgatzh/internal/examples/example1"
	"github.com/go-zwbc/pkgatzh/internal/examples/example2"
	"github.com/yyle88/neatjson/neatjsons"
)

// TestExample tests both example1 and example2 package information extraction
// Verifies that distinct packages can retrieve path and name data on its own
//
// TestExample 测试 example1 和 example2 两个包的信息提取
// 验证不同包可以独立获取路径和名称数据
func TestExample(t *testing.T) {
	// Log example1 package information (记录 example1 包信息)
	t.Log(neatjsons.S(example1.V位置信息))
	// Log example2 package information (记录 example2 包信息)
	t.Log(neatjsons.S(example2.V位置信息))
}
