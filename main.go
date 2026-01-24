package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Calvados-1997/monkeyInterpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s, Welcome to Monkey language World!\n", user.Username)
	fmt.Printf("Type command you want to execute\n")
	repl.Start(os.Stdin, os.Stdout)
}
