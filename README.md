# env

A simple environment variable manager for Go.

To use you import it and call the NewEnv function.

```go
package main

import (
	"fmt"
	"github.com/nathanbcrocker/env"
)

func main() {
	// Create a new environment manager
	e := env.NewEnv()

	// Get an environment variable
	port, ok := e.Get("PORT")
	if ok {
		fmt.Println("PORT:", port.Value)
	} else {
		fmt.Println("PORT not set")
	}
}
```

It will automatically load environment variables from a `.env` file in the current directory. 
Additionally, when e.Get() is called, it will attempt to load the environment variable from 
the system environment variables, if it was not in the `.env` file.