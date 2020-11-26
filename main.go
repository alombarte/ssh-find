package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var appName = "ssh-find"
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

func main() {

	flag.Usage = func() {
		fmt.Printf(`
%s v%s (%s)

Usage: %s <username>
Finds ssh public keys by username
`, appName, version, date, appName)
		flag.PrintDefaults()
	}

	flag.Parse()
	username := flag.Arg(0)

	if len(username) < 1 {
		fmt.Printf(`
%s v%s (%s)

Usage: %s <username>
Finds ssh public keys by username
`, appName, version, date, appName)
		os.Exit(1)
	}

	fmt.Println(GetSSHKeys(username))
}

// GetSSHKeys from public git providers
func GetSSHKeys(username string) string {

	providers := map[string]string{
		"Github": "https://github.com/%s.keys",
		"Gitlab": "https://gitlab.com/%s.keys",
	}

	foundKeys := ""
	for provider, url := range providers {
		url = fmt.Sprintf(url, username)
		keys := getURLContent(url)
		if len(keys) > 0 {
			foundKeys += fmt.Sprintf("\n# %s (%s):\n%s", username, provider, keys)
		}
	}

	return foundKeys

}

func getURLContent(url string) string {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return fmt.Sprintf("%s", data)
	}

	return ""
}
