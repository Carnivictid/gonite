/* Gonite main package.
=============================
*make an arg: gonite.exe -i jsonname to run custom json
chatmansam.com/gonite
This downloads the exe
Exe runs, gets chatmansam.com/gonite/json.json to get updated json
downloads the files and installs them
Spits out a report of what happened
*/

package main

import (
	"log"
	"./bin"
)

func main() {
	log.Println("Test log.")
	gonite.TestFunc()
}
