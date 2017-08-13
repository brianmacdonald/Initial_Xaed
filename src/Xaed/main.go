package main

import (
	"Xaed/repltwo"
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s. This is the REPL.\n",
		user.Username)
	repltwo.Start()
}
