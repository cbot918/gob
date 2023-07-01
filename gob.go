package main

import (
	"fmt"

	"github.com/cbot918/tools/pkg/ygo/scanner"
	"github.com/cbot918/tools/pkg/ygo/yutil"
	"github.com/cbot918/tools/pkg/ygo/yutil/command"
)

var (
	log  = fmt.Println
	logf = fmt.Printf
)

func main() {
	// c := command.NewCmd()

	parsedSource := scanner.NewScanner("gob.json").GetObjArr()
	yutil.LogJson(parsedSource)

	y := command.NewYexec("testp")
	y.ExecObj(parsedSource)

}

// func Logj(target []map[string]string) {
// 	b, err := json.MarshalIndent(target, "", "  ")
// 	if err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	for _, item := range b {
// 		fmt.Print(string(item))
// 	}

// }
