package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, _ string) error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return err
}
