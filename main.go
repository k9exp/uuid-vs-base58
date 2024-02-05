package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		id, err := New(21)

		if err != nil {
			panic(err)
		}

		time.Sleep(500 * time.Millisecond)

		fmt.Println(id)
	}
}
