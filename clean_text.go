package main

import "strings"

func cleanInput(t string) []string {
	o := strings.ToLower(t)
	s := strings.Fields(o)
	return s
}
