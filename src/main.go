// main package for gonite

package main

import (
	"log"
	"./bin"
)

func main() {
	// pkgs is an array of structs
	// Exe `json:"exe"`, Url `json:"url"`, Flg `json:"flg"`
	pkgs := gonite.GetPkgsFromJson()
	for _, p := range pkgs {
		log.Printf("Downloading: %v", p.Exe)
		gonite.DownloadFile(p.Exe, p.Url)
		log.Printf("Download Complete")
	}
}
