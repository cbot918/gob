package yutil

import (
	"encoding/json"
	"fmt"
	"log"
)

func Checke(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Checkt(obj interface{}) {
	fmt.Printf("\n type: %T \n", obj)
}

func Logg(obj interface{}) {
	fmt.Println(obj)
}

func TrimS(s string, n int) string {
	ns := ""
	if n >= 0 {
		ns = s[n:len(s)]
	}
	if n <= -1 {
		ns = s[0 : len(s)+n]
	}
	return ns
}

func LogJson(target []map[string]string) {
	b, err := json.MarshalIndent(target, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, item := range b {
		fmt.Print(string(item))
	}

}
