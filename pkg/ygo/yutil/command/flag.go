package command

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	projName string
	fileName string
}

func NewCmd() *Cmd {
	c := new(Cmd)
	// c.fileName = "config.json"
	c.ParseArg()
	c.ParseFlag()
	return c
}

func (c *Cmd) ParseArg() {
	c.projName = os.Args[1:][0]
}

func (c *Cmd) ParseFlag() {
	var nFlag = flag.String("f", "config.json", "example: gob myproj -f myconfig.json")
	flag.Parse()
	fmt.Println(" ***** nFlag =", *nFlag)
	c.fileName = *nFlag
}

func (c *Cmd) FileName() string {
	return c.fileName
}

func (c *Cmd) ProjName() string {
	return c.projName
}
