package hosts

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var hostsPath = "/etc/hosts"
var tag = "#he-added-this"

// Write an entry to the hosts file
func Add(ip string, hostname string) error {

	file, err := ioutil.ReadFile(hostsPath)
	if err != nil {
		return err
	}

	contents := string(file)
	rows := strings.Split(contents, "\n")

	newLine := fmt.Sprintf("%s %s %s\n", ip, hostname, tag)
	newRows := append(rows, newLine)

	newContents := strings.Join(newRows, "\n")

	err = ioutil.WriteFile(hostsPath, []byte(newContents), 0x644)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
