package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal("no resource is mentioned \nExample command : ls pods")
	}

	if argsWithoutProg[0] != "pods" {
		log.Fatal("unknown resource kind")
	}

	fmt.Println("lets print the pods list")
}
