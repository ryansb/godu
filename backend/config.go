package backend

import (
	"code.google.com/p/gcfg"
	"fmt"
	"github.com/op/go-logging"
	stdlog "log"
	"os"
)

var log = logging.MustGetLogger("godu")

func init() {
	logBackend := logging.NewLogBackend(os.Stderr, "[godu] ", stdlog.LstdFlags|stdlog.Lshortfile)
	logBackend.Color = true
	logging.SetBackend(logBackend)
	logging.SetLevel(logging.DEBUG, "")
}

type Config struct {
	Admin struct {
		// but log when you *would* have run them
		BackEnds []BackEnd
		// how many concurrent jobs can be run by the scheduler
		MaxWorkers int64
		// maxumim length a job can run (seconds)
		MaxTime int64
		// defaults to /tmp/godu_logs
		LogDir int64
		// defaults to /tmp/godu.log
		LogFile int64
		// run no jobs at all for the present time
		// but log when you *would* have run them
		SuspendAll bool
	}
}

func ReadConfig(loc string) (Config, error) {
	var conf Config
	err := gcfg.ReadFileInto(&conf, loc)
	if err != nil {
		fmt.Println("Couldn't read file.")
		return conf, err
	}
	return conf, nil
}
