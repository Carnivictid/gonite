// learning JSON

package gonite

import (
	"encoding/json"
	"log"
	"io/ioutil"
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
