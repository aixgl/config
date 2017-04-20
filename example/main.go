package main

import (
	"fmt"
	"github.com/aixgl/config"
	"os"
)

func main() {

	//Get absolute file path
	filepath, _ := os.Getwd()
	filepath += "/config.env"
	conf := config.C("G", filepath)
	confE := config.C()
	fmt.Println(confE.Get(), filepath)
	fmt.Println(conf.GetAll(), filepath)

	fmt.Println(confE.Get())
	debuglog, _ := confE.Get("DEBUG.log").(string)
	fmt.Println(debuglog)
	fmt.Println(confE.Get("DEBUG", "log"))
	fmt.Println(confE.Get("DEBUG", "cms"))
	debug, _ := confE.Get("DEBUG").(map[string]string)
	fmt.Println(debug)
}
