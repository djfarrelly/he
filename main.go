package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/codegangsta/cli"
)

var hostsPath = "/etc/hosts"

func main() {

	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "a hosts file editing tool"
	app.Version = "0.1"

	app.Author = "Dan Farrelly"
	app.Email = "https://github.com/djfarrelly"

	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "add an entry",
			Action: func(c *cli.Context) {

				ip := c.Args().Get(0)
				hostname := c.Args().Get(1)

				fmt.Printf("adding entry \"%s %s\"\n", ip, hostname)

				file, err := ioutil.ReadFile(hostsPath)
				if err != nil {
					log.Fatal(err)
				}

				contents := string(file)
				rows := strings.Split(contents, "\n")

				// for idx, element := range rows {
				// 	if idx < 10 {
				// 		fmt.Println(element)
				// 	}
				// }

				newLine := fmt.Sprintf("%s %s #he-added-this\n", ip, hostname)
				newRows := append(rows, newLine)

				newContents := strings.Join(newRows, "\n")

				fmt.Println(newContents);
			
				err = ioutil.WriteFile(hostsPath, []byte(newContents), 0x644)
				if err != nil {
					log.Fatal(err)
				}

			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
