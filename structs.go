package main

type Config struct {
	Jobs      []JobSimple `json:"jobs"`
	Admin     AdminBlock `json:"admin"`
}

type AdminBlock struct {
	// how many concurrent jobs can be run by the scheduler
	MaxWorkers    int64 `json:"maxworkers"`
	// maxumim length a job can run (seconds)
	MaxTime       int64 `json:"maxtime,omitempty"`
	// defaults to /tmp/godu_logs
	LogDir        int64 `json:"logdir,omitempty"`
	// defaults to /tmp/godu.log
	LogFile       int64 `json:"logfile,omitempty"`
	// run no jobs at all for the present time
	// but log when you *would* have run them
	SuspendAll    bool  `json:"suspendall,omitempty"`
	// 0-5, where 0 is "pedantically accurate" and
	// 5 is "when you get a chance"
	Granularity   int   `json:"granularity"`
}

type JobSimple struct {
	// a nickname for the job, totally optional
	ShortName   string `json:"name,omitempty"`
	// the block for selecting the frequency of the job
	Freq        FrequencyBlock `json:"frequency"`
	// fully qualified path for the executable
	FQExec      string `json:"fullpath"`
	// any args that need to be passed to it
	Args        string `json:"args,omitempty"`
	// true if the job should not run
	Suspend     bool `json:"suspended,omitempty"`
}

type FrequencyBlock struct {
	Second       int64 `json:sec,omitempty`
	IntSecond    int64 `json:intsec,omitempty`
	Minute       int64 `json:min,omitempty`
	IntMinute    int64 `json:intmin,omitempty`
	Hour         int64 `json:hour,omitempty`
	IntHour      int64 `json:inthour,omitempty`
	Day          int64 `json:day,omitempty`
	IntDay       int64 `json:intday,omitempty`
	WDay         int64 `json:weekday,omitempty`
	IntWDay      int64 `json:intweekday,omitempty`
	Month        int64 `json:month,omitempty`
	IntMonth     int64 `json:intmonth,omitempty`
}
