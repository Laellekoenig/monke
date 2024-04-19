package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Laellekoenig/monke/internal/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Monke\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
