package pkgpathzh

import (
	"github.com/yyle88/runpath"
	"github.com/yyle88/syntaxgo"
)

type T路径和名 struct {
	P路径 string
	N包名 string
}

func NewT路径和名() *T路径和名 {
	return &T路径和名{
		P路径: runpath.PARENT.Skip(1),
		N包名: syntaxgo.GetPkgName(runpath.Skip(1)),
	}
}
