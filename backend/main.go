package main

import (
	"fmt"

	"github.com/vprates-22/pratesBet/internal/db"
)

func main() {
	err := db.New()
	if err != nil {
		fmt.Println(err)
	}
}
