package main

import (
	//"encoding/json"
	"flag"
	"fmt"
	"os"
)

// args for the job being added
var JobArguments string

const (
	configCmd      = "config"
	makeConfigCmd  = "mkconfig"
	helpCmd        = "help"
	addCmd         = "add"
	delCmd         = "del"
	configLoc      = "godu.gocfg"
	version        = "godu version 0.1\nhttps://github.com/ryansb/godu "
	helpText       = "godu is a job scheduling application that runs in " +
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
	fs := flag.NewFlagSet("godu", flag.ExitOnError)
	const (
		usageArgs = "Arguments to pass to the executable when it runs"
	)
	fs.StringVar(&JobArguments, "args", "", usageArgs)
	fs.StringVar(&JobArguments, "a", "", usageArgs)

	if len(os.Args) <= 1 {
		fmt.Println("No args. Wat?")
		return
	} else if len(os.Args) == 2 {
		if string(os.Args[1]) == "--help" ||
			string(os.Args[1]) == "-h" {
			fmt.Println(helpText)
			return
		} else if string(os.Args[1]) == "--version" ||
			string(os.Args[1]) == "-v" {
			fmt.Println(version)
			return
		}
	} else {
		fs.Parse(os.Args[2:])
	}

	switch os.Args[1] {
	case makeConfigCmd:
		fmt.Println("make config")
	case addCmd:
		fmt.Println("Add unimplemented.")
	case delCmd:
		fmt.Println("Del unimplemented.")
	case helpCmd:
		fmt.Println(helpText)
	default:
		fmt.Println("Subcommand not found.")
	}
}
