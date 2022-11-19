package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/theFr3Y/color"
)

var logo = color.LightCyan + `
  ________         ________        __  .___        _____       
 /  _____/  ____  /  _____/  _____/  |_|   | _____/ ____\____  
/   \  ___ /  _ \/   \  ____/ __ \   __\   |/    \   __\/  _ \ 
\    \_\  (  <_> )    \_\  \  ___/|  | |   |   |  \  | (  <_> )
 \______  /\____/ \______  /\___  >__| |___|___|  /__|  \____/ 
        \/               \/     \/              \/             
`

func main() {
	Clear()
	var input string
	fmt.Println(logo)
	fmt.Printf(color.Reset + "type your target address example\"google.com\": " + color.LightYellow)
	fmt.Scanln(&input)

	resp, err := http.Get("http://ip-api.com/json/" + input)
	if err != nil {
		fmt.Println(" ")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println(" ")
	}
	for k, v := range result {
		fmt.Printf(color.LightPurple+"%v -> "+color.Green+"%v\n "+color.Reset, k, v)
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func Clear() {
	system := runtime.GOOS
	if system == "windows" {
		console := exec.Command("cmd", "/c", "cls")
		console.Stdout = os.Stdout
		console.Run()
	} else {
		console := exec.Command("clear")
		console.Stdout = os.Stdout
		console.Run()
	}
}
