package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner:= bufio.NewScanner(os.Stdin)

	fmt.Print("domain,hasMX,hasSPF,")
}