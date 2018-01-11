package gonite

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
	"io"
	"os"
)

type Pkg struct {
    Exe string `json:"exe"`
    Url string `json:"url"`
    Flg string `json:"flg"`
}


//a little better practice than inline strings
const exeDir string = "C:/temp/"
const fExtension string = ".exe"


func GetPkgsFromJson() []Pkg {
	raw, err := ioutil.ReadFile("./bin/squires.json")
	if err != nil {
		log.Fatal(err)
	}
	var c []Pkg
	json.Unmarshal(raw, &c)
	return c
}

//Here were going to pass it the package and a channel to report to once its downloaded.
func DownloadFile(pack Pkg, downloads chan Pkg) {
	//first time I ran it, crashed because the file path didn't exist.
	//so let's check if dir exists, if not, create it.
	if _, err := os.Stat(exeDir); os.IsNotExist(err) {
    os.Mkdir(exeDir, 0777)
  }
	// create a blank file named "filename"
	out, err := os.Create(exeDir + pack.Exe + fExtension)
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
	downloads <- pack
}
