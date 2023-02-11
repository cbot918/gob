package gob

import (
	"fmt"
	"os"

	"github.com/cbot918/liby/cmdy"
	u "github.com/cbot918/liby/util"
)


type Gob struct {}

func New() *Gob{
	g := new(Gob)

	fmt.Println("hello this is gob ( go-builder )")

	return g
}
func (g *Gob) Run(){
	c := cmdy.New()
	args := os.Args
	projectName := ""
	switch args[1] {
	case "init":{
		if args[2] == "."{
			u.Logg("in .")
		}else {
			projectName = args[2]
			// u.Logg(fmt.Sprintf("name: %s", projectName))

			c.Run([]string{"git clone -b go https://github.com/cbot918/template"})
			c.Run([]string{fmt.Sprintf("mv template %s", projectName)})

		}
	}
	case "gitc":{
		if args[2] == "."{
			c.Run([]string{
				"rm -rf .git && git init",
				"git checkout -b main",
				fmt.Sprintf("echo \"# %s\" > README.md", projectName ),
				"git add .",
				"git commit -m \"init project\"",
				"gh repo create $name --public --push --source .",
			})
		}else{
			fmt.Println("後面+個點")
			return
		}
	
	}
	}
}