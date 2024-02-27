package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gen2brain/beeep"
)

var appName = flag.String("name", "tee-notify", "name of the app")
var search = flag.String("search", "restarted", "string to search for")

// extend because ...
const maxCapacity = 64 * 1000 * 2

func main() {
	flag.Parse()

	for {
		err := scan()
		if err == nil {
			break
		}
		fmt.Println("Scanner fail: ", err)
		fmt.Println("---- Sending notification about continue-scanning ---")
		err = beeep.Notify(*appName, *appName+" scan error. Continue scanning on next line", "")
		if err != nil {
			panic(err)
		}
	}
}

func scan() error {

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	i := 0
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Println(i, txt)
		i++
		if strings.Contains(txt, *search) {
			fmt.Println("---- Sending notification ---")
			err := beeep.Notify(*appName, *appName+" restarted", "")
			if err != nil {
				panic(err)
			}
		}
	}

	return scanner.Err()

}
