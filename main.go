package main

import (
	"fmt"
	"os"
	"os/user"
	"singo/repl"
)

func main() {
	user, err:= user.Current()
	if err!=nil{
		panic(err)
	} 
	fmt.Printf("Hello %s! This is the Singo programming language!\n",user.Username)
	fmt.Printf("Feel Free to type in Commands\n")
	repl.Start(os.Stdin, os.Stdout)
}