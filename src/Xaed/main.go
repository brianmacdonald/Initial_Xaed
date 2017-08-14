package main

import (
	"fmt"
	"os"
	"os/user"
	"Xaed/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Xaed programming language!\n",
		user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
