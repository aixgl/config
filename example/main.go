package main

import (
	"fmt"
	"github.com/aixgl/config"
)

func main() {

	//Get absolute file path
	conf := config.C("G", "config.env")
	confE := config.C()
	fmt.Println(confE.Get())
	fmt.Println(conf.GetAll())

	fmt.Println(confE.Get())
	debuglog, _ := confE.Get("DEBUG.log").(string)
	fmt.Println(debuglog)
	fmt.Println(confE.Get("DEBUG", "log"))
	fmt.Println(confE.Get("DEBUG", "cms"))
	debug, _ := confE.Get("DEBUG").(map[string]string)
	fmt.Println(debug)
}
