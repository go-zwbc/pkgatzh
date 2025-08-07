package example1_test

import (
	"testing"

	"github.com/go-zwbc/pkgpathzh/internal/examples/example1"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestExample(t *testing.T) {
	t.Log(neatjsons.S(example1.V路径和名))
}
