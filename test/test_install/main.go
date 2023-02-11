package main

import (
	"github.com/cbot918/liby/cmdy"
)

func main(){

	c := cmdy.New()
	
	script := []string{"git clone -b go https://github.com/cbot918/template"}
	
	c.Run(script)
}