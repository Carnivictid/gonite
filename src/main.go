// main package for gonite

package main

import (
	"log"
	"./bin"
	"os/exec"
	"os"
)

func main() {
	// pkgs is an array of structs
	// Exe `json:"exe"`, Url `json:"url"`, Flg `json:"flg"`
	pkgs := gonite.GetPkgsFromJson()
	
	for _, p := range pkgs {
		log.Printf("Installing: %v", p.Exe)
		gonite.DownloadFile(p.Exe, p.Url)
		RunExe(p.Exe, p.Flg)
		log.Printf("%v has finished installing.", p.Exe)
	}
}

func RunExe(exe string, flg string) {
	os.Chdir("C:/temp/") // files are saved in temp.
	exec.Command(exe + ".exe ", flg).Run()
}
