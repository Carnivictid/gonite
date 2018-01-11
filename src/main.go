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

	//create a buffered channel the size of our json
	downloads := make(chan gonite.Pkg, len(pkgs))

	for _, p := range pkgs {
		log.Printf("Starting Download for: %v", p.Exe)
		//spin off each one so downloads happen concurrently
		go gonite.DownloadFile(p, downloads)
	}

	//As values come into the downloads channel, they start downloading.
	//This blocks until a value is received, so installs happen
	//in the order in which they finished downloading.
	for i := range downloads{
		RunExe(i.Exe, i.Flg)
		log.Printf("%v has finished installing.", i.Exe)
	}
}

func RunExe(exe string, flg string) {
	os.Chdir("C:/temp/") // files are saved in temp.
	exec.Command(exe + ".exe ", flg).Run()
}
