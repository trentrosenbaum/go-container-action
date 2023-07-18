package main

import (
	"fmt"
	"github.com/actions-go/toolkit/core"
	"os"
)

func main() {
	myInput := os.Getenv("INPUT_MYINPUT")
	core.Debug(fmt.Sprintf("Input set to %s", myInput))

	core.Info("Creating myOutput value")
	myOutput := fmt.Sprintf("Hello %s", myInput)

	core.SetOutput("myOutput", myOutput)
	core.Debug(fmt.Sprintf("Output set to %s", myOutput))
}
