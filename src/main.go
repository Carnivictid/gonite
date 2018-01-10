// main package for gonite

package main

import (
	"log"
	"./bin"
)

func main() {
	pkgs := gonite.GetPkgsFromJson()
	for _, p := range pkgs {
		log.Printf("Exe: %v | Url: %v | Flg: %v\n", p.Exe, p.Url, p.Flg)
	}
}
