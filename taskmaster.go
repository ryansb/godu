package main

import (
	"fmt"
	"sync"
	"time"
	gocal "github.com/mish33/go-calendar"
)

/*
func runWithTimeout(callable interface{}, timeout int64) {
	c := make(chan string)
	go callable(&c)
	select {
		case m := <-c:
			handle(m)
		case <-time.After(5 * time.Minute):
			fmt.Println("timed out")
	}
}*/

type Job struct {
	Guard sync.Mutex
	Msg JobMsg
}

func (job Job) Run(abort *chan bool) (){
	t := time.Now()
	first := time.Date(t.Year(), 0, 0, 0, 0, 0, 0, t.Location())
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	if job.HappensToday() {

	}
	interval = time.Since(midnight) % job.GetIntraDayRotation()
	var interval time.Duration
	if job.GetIntraDayRotation() == 0 {
	} else if job.Msg.GetFrequency().GetWeekday() == -1 {
		interval = time.Since(first) % job.GetInterDayRotation()
	} else {
		sunday := time.Date(t.Year(), 0, 0, 0, 0, 0, 0, t.Location())
		interval = time.Since(sunday) % job.GetInterDayRotation()
	}

	gotlock := make(chan bool)
	go func () {
		job.Guard.Lock()
		gotlock <- true
	}()
	select {
	case <- gotlock:
		go func() {
			defer job.Guard.Unlock()
			timer := time.After(interval)
			for {
				select {
				case <- timer:
					// do job
				case <- *abort:
					// if, before we're done, the app exits, then GTFO.
					return
				}
			}
		}()
	case <- time.After(time.Second * 60):
		return
	}
}

func (job Job) GetRotation() (rotate time.Duration) {
	freq := job.Msg.GetFrequency()
	if freq.GetSecond() <= 0 {
		rotate = rotate + time.Second * time.Duration(freq.GetSecond())
	}
	if freq.GetMinute() <= 0 {
		rotate = rotate + time.Minute * time.Duration(freq.GetMinute())
	}
	if freq.GetHour() <= 0 {
		rotate = rotate + time.Hour * time.Duration(freq.GetHour())
	}
	return
}

func (job Job) HappensToday() (bool) {
	freq := job.Msg.GetFrequency()
	if freq.GetWeekday == gocal.Weekday(t.Year(), t.Month(), t.Day()) {
		return true
	}
	if freq.GetWeekday() == -1{
		// normal interval
	} else {
	}
}
