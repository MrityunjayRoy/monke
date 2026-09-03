package main

import (
	"fmt"
	"monke/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("monke together stronk")
	fmt.Printf("want banana: %s\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
