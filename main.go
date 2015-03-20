package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/codegangsta/cli"
	"github.com/djfarrelly/he/hosts"
)

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

				err := hosts.Add(ip, hostname)
				if err != nil {
					log.Fatal(err)
				}

			},
		},
		{
			Name:  "rm",
			Usage: "remove an entry",
			Action: func(c *cli.Context) {

				fmt.Printf("removing entry \"%s %s\"\n")

			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
