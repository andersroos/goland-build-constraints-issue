package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func main() {
	var db sqlx.DB
	fmt.Printf("var type: %T\n", db)
}
