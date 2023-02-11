package main

import "github.com/cbot918/liby/cmdy"

func main(){
	c := cmdy.New()
	script1 := []string{"git clone -b go https://github.com/cbot918/template"}
	script2 := []string{"rm -rf template"}
	c.Run(script1)
	c.Run(script2)
	
	

}