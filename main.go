// main.go
package main

import (
	"fmt"
)

func main() {
	// Create a new Env instance
	e := NewEnv()

	// Get and print the "PORT" value
	port, ok := e.Get("PORT")
	if ok {
		fmt.Println("PORT:", port.Value)
	} else {
		fmt.Println("PORT not set")
	}
}
