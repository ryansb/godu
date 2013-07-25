package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ryansb/godu/backend"
	"os"
)

// args for the job being added
var JobArguments string

const (
	configLoc = "godu.gocfg"
	helpText  = "godu is a job scheduling application that runs in " +
		"the background and repeats infrequent jobs.\n" +
		"Usage:\ngodu add <executable> every 10 minutes\n" +
		"# provide arguments to pass to the executable when it runs\n" +
		"godu add <executable> every hour -args '-u ryansb -a'\n" +
		"godu add <executable> every other week\n" +
		"# delete the first job (as shown in `godu list`) for a specific " +
		"executable\n" +
		"godu del 1 <executable>\n" +
		"# show all executables that are scheduled\n" +
		"godu list \n" +
		"# show jobs scheduled for a specific executable\n" +
		"godu list <executable>"
)

func main() {
	app := cli.NewApp()
	app.Name = "godu"
	app.Version = "0.1"
	app.Usage = helpText
	app.Commands = []cli.Command{
		{
			Name:      "add",
			ShortName: "a",
			Action: func(c *cli.Context) {
				fmt.Println("Add unimplemented.")
				fmt.Println(c.Args)
			},
		},
		{
			Name:      "delete",
			ShortName: "rm",
			Action: func(c *cli.Context) {
				fmt.Println("Delete unimplemented.")
			},
		},
		{
			Name:      "list",
			ShortName: "ls",
			Action: func(c *cli.Context) {
				fmt.Println("List unimplemented.")
			},
		},
		{
			Name:      "config",
			ShortName: "c",
			Action: func(c *cli.Context) {
				fmt.Println("Config check.")
				_, err := backend.ReadConfig(configLoc)
				if err != nil {
					fmt.Println("Something is wrong with the config.")
					fmt.Println(err)
				}
				fmt.Println("Config is valid.")

				job := backend.Job{}
				fmt.Println(job.GetRotation())
			},
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{"args", "", "Arguments to pass to the executable"},
	}
	app.Run(os.Args)
}
