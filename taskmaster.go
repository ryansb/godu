package main

import (
	"sync"
	"time"
)

type Job struct {
	Guard sync.Mutex
	Msg JobMsg
}

func (job Job) Run(abort *chan bool) (){
	t := time.Now()
	if ! job.HappensOn(&t) { return }

	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	interval := time.Since(midnight) % job.GetRotation()

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
	case <- time.After(time.Second * 30):
		return
	}
}

func (job Job) GetRotation() (rotate time.Duration) {
	freq := job.Msg.GetFrequency()
	if freq.GetSecond() >= 0 {
		rotate += time.Second * time.Duration(freq.GetSecond())
	}
	if freq.GetMinute() >= 0 {
		rotate += time.Minute * time.Duration(freq.GetMinute())
	}
	if freq.GetHour() >= 0 {
		rotate += time.Hour * time.Duration(freq.GetHour())
	}
	return
}

func (job Job) HappensOn(t *time.Time) (bool) {
	freq := job.Msg.GetFrequency()

	return int8(freq.GetWeekday()) == int8(t.Weekday()) ||
	freq.GetWeekday() == -1 &&
	freq.GetMonth() == int32(t.Month()) &&
	freq.GetDay() == int32(t.Day())
}
