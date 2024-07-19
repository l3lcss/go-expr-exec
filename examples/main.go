package main

import (
	"fmt"
	"log"

	"github.com/l3lcss/go-expr-exec"
)

func main() {
	result, err := expr.Execute[int, int](1, expr.RelationalOperatorEqualTo, 1)
	if err != nil {
		log.Fatalf("Error %s", err.Error())
	}

	fmt.Printf("result: %v\n", result)
}
