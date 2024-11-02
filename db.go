package main

import (
	"my-api/migration"
	"my-api/seed"
	"os"
	"fmt"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		if args[1] == "migrate" {
			fmt.Println("Migrating database...")
			migration.Run()
		} else if args[1] == "seed" {
			fmt.Println("Seeding database...")
			seed.Run()
		}
	}
}


