package gonite

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os/exec"
	"log"
	"io"
	"os"
)

type Pkg struct {
	// Contains the filename, download location, and any cmd flags
	// needed to make sure the install runs correctly.
    Exe string `json:"exe"`
    Url string `json:"url"`
    Flg string `json:"flg"`
}

// constants containing file structure information
const exeDir string = "C:/temp/"
const fileExt string = ".exe"


func GetPkgsFromJson() []Pkg {
	// this function returns an array of Pkg structs.
	raw, err := ioutil.ReadFile("./bin/squires.json")
	if err != nil {
		log.Fatal(err) // TODO add code to download json from git repo
	}

	var c []Pkg
	err = json.Unmarshal(raw, &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// Pass it the package and a channel to report to once its downloaded.
func DownloadFile(pack Pkg, downloads chan Pkg) {
	// check if download dir exists, create it if not
	if _, err := os.Stat(exeDir); os.IsNotExist(err) {
		os.Mkdir(exeDir, 0777)
	}
	// create a blank file named "filename"
	out, err := os.Create(exeDir + pack.Exe + fileExt)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close() // Close the file when we finish

	// Get the data from the url
	resp, err := http.Get(pack.Url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() // Close the response when we finish

	// Write the body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//send the package to downloads channel, indicating its done dowloading
	defer log.Printf("%v has finished downloading.", pack.Exe)
	downloads <- pack
}

func RunExe(exe string, flg string) {
	os.Chdir(exeDir) // TODO gotta be a better way to do this

	err := exec.Command(exe + fileExt, flg).Run()
	if err != nil {
		log.Fatal(err)
	}
}
