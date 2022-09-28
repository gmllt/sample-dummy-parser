package main

import (
	"fmt"
)

func main() {
	log := FromStdin()
	fmt.Println(log.Correct())

}
