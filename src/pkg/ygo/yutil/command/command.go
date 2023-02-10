package command

import (
	"gob/src/pkg/ygo/yutil"
	"os/exec"
)

type Yexec struct {
	projName string
}

func NewYexec(projName string) *Yexec {
	y := new(Yexec)
	y.projName = projName
	return y
}

func (y *Yexec) Exec(cmda string, cmdb string) {
	cmd := exec.Command(cmda, cmdb)
	err := cmd.Run()
	yutil.Checke(err)
}

func (y *Yexec) ExecObj(objArr []map[string]string) {
	yutil.Logg("projname: " + y.projName)
	// y.projName = "hhh"
	y.Exec("mkdir", y.projName)
	for _, obj := range objArr {
		if obj["type"] == "file" {
			y.Exec("touch", y.projName+"/"+obj["name"])
		} else if obj["type"] == "folder" {
			y.Exec("mkdir", y.projName+"/"+obj["name"])
		}
		yutil.Logg(obj)
	}
}
