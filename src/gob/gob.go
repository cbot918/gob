package gob

import (
	"fmt"
	"os"

	"github.com/cbot918/liby/cmdy"
	u "github.com/cbot918/liby/util"
)


type Gob struct {
	GhName string
	GhUrl string
}

func New() *Gob{
	g := new(Gob)

	g.GhName = "gh_2.23.0_linux_amd64"
	g.GhUrl = "https://github.com/cli/cli/releases/download/v2.23.0/gh_2.23.0_linux_amd64.tar.gz"

	return g
}
func (g *Gob) Run(){
	c := cmdy.New()
	args := os.Args
	projectName := ""

	if len(args) == 1{ fmt.Println("gob help page")} else {
		switch args[1] {
			case "cmd":{
				if len(args) == 2 { fmt.Println("gob cmd help page") } else {
					u.Logg("in cmd")
					c.Run([]string{
						args[2],
					})
				}
			}
		
			case "init":{
				if len(args) == 2 { fmt.Println("gob init help page") } else {
					if args[2] == "."{
						c.Run([]string{
							"git clone -b go https://github.com/cbot918/template",
							"mv template/* ./",
							"rm -rf template",
						})
						
					}else {
						projectName = args[2]
						c.Run([]string{
							"git clone -b go https://github.com/cbot918/template",
							fmt.Sprintf("mv template %s", projectName),
						})
					}
				}
			}

			case "gitc":{
				if len(args) == 2 { fmt.Println("gob gitc help page")} else{
					if args[2] == "."{
						u.Logg("in gitc .")
						c.Run([]string{
							"rm -rf .git && git init",
							"git checkout -b main",
							// fmt.Sprintf("echo \"# %s\" > README.md", projectName ),
							"git add .",
							"git commit -m 'init project'",
							"echo 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 | gh auth login --with-token",
							"gh repo create --public --push --source .",
						})
					}else{
						fmt.Println("後面+個點")
						return
					}
				}
			}
				// case "install": {
			// 	if args[2] == "docker"{
			// 		c.Run([]string{
			// 			"apt update && apt install curl -y",
			// 			fmt.Sprintf("curl -OL %s && tar -xvf %s.tar.gz",g.GhUrl, g.GhName),
			// 			fmt.Sprintf("mkdir -p /usr/local/bin && cp %s/bin/gh /usr/local/bin ",g.GhName),
			// 			fmt.Sprintf("rm -r %s && rm %s.tar.gz",g.GhName,g.GhName ),
			// 			"git config --global user.name 'cbot918'",
			// 			"git config --global user.email 'cbot918@gmail.com'",
			// 		})
			// 	}
			// 	if args[2] == "vm" {
			// 		c.Run([]string{
			// 			"sudo apt update && sudo apt install curl",
			// 			fmt.Sprintf("curl -OL %s && tar -xvf %s.tar.gz",g.GhUrl, g.GhName),
			// 			fmt.Sprintf("mkdir -p /usr/local/bin && sudo cp %s/bin/gh /usr/local/bin ",g.GhName),
			// 			fmt.Sprintf("rm -rf %s && rm %s.tar.gz",g.GhName,g.GhName),
			// 			"git config --global user.name 'cbot918'",
			// 			"git config --global user.email 'cbot918@gmail.com'",
			// 		})
			// 	}
		
			// }
			}
	}

}