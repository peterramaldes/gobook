// Echo3 show the args from the command line
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
}
