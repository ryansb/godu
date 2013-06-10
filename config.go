package main

import (
	"fmt"
	"code.google.com/p/gcfg"
)

type Config struct {
	Admin struct {
		// how many concurrent jobs can be run by the scheduler
		MaxWorkers    int64
		// maxumim length a job can run (seconds)
		MaxTime       int64
		// defaults to /tmp/godu_logs
		LogDir        int64
		// defaults to /tmp/godu.log
		LogFile       int64
		// run no jobs at all for the present time
		// but log when you *would* have run them
		SuspendAll    bool
		// 0-5, where 0 is "pedantically accurate" and
		// 5 is "when you get a chance"
		Granularity   int
	}
}

func readConfig(loc string) (Config, error) {
	var conf Config
	err := gcfg.ReadFileInto(&conf, loc)
	if err != nil {
		fmt.Println("Couldn't read file.")
		return conf, err
	}
	return conf, nil
}
