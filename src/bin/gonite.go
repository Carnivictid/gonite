// learning JSON

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

func GetPkgsFromJson() []Pkg {
	raw, err := ioutil.ReadFile("./bin/squires.json")
	if err != nil {
		log.Fatal(err)
	}
	var c []Pkg
	json.Unmarshal(raw, &c)
	return c
}

func DownloadFile(filename string, url string) error {
	// create a blank file named "filename"
	out, err := os.Create("C:/temp/" + filename + ".exe")
	if err != nil {
		return err
	}
	defer out.Close() // Close the file when we finish
	
	// Get the data from the url
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // Close the response when we finish
	
	// Write the body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
