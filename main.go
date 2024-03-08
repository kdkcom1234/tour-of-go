package main

import (
	"fmt"
	"time"

	"github.com/kdkcom1234/tour-of-go/welcome"
)

func main() {
	welcome.Hello()
	welcome.Sandbox()
	welcome.Imports()
	welcome.ExportedNamesGo()
	welcome.Functions()
	welcome.FunctionsContinued()
	for {
		fmt.Println(time.Now())
		time.Sleep(1 * time.Second)
	}
}
