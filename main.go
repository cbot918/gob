package main

import (
	"gob/src/pkg/ygo/scanner"
	"gob/src/pkg/ygo/yutil"
	"gob/src/pkg/ygo/yutil/command"
)

func main() {
	c := command.NewCmd()

	s := scanner.NewScanner(c.FileName())
	yutil.LogJson(s.GetObjArr())

	y := command.NewYexec(c.ProjName())
	y.ExecObj(s.GetObjArr())

}
