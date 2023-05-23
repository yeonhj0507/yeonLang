package main

import (
	"fmt"
	"os"
	"os/user"
	"yeonLang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the NewLang programming language!\n", user.Username)
	fmt.Printf("Feel Free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}