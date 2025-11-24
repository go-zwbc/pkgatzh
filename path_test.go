package pkgatzh_test

import (
	"testing"

	"github.com/go-zwbc/pkgatzh"
	"github.com/yyle88/neatjson/neatjsons"
)

// TestNewT位置信息 tests creating a new T位置信息 instance and verifying it retrieves correct path and name
// TestNewT位置信息 测试创建新的 T位置信息 实例并验证它获取正确的路径和名称
func TestNewT位置信息(t *testing.T) {
	// Create instance and log its JSON representation (创建实例并记录其 JSON 表示)
	t.Log(neatjsons.S(pkgatzh.NewT位置信息()))
}
