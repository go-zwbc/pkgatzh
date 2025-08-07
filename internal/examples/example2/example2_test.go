package example2_test

import (
	"testing"

	"github.com/go-zwbc/pkgpathzh/internal/examples/example1"
	"github.com/go-zwbc/pkgpathzh/internal/examples/example2"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestExample(t *testing.T) {
	t.Log(neatjsons.S(example1.V路径和名))
	t.Log(neatjsons.S(example2.V路径和名))
}
