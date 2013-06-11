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
	if ! job.HappensOn(&t) {
		log.Debug("Job", job.Msg.GetName(), "doesn't happen today. Bailing out.")
		return
	}

	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	interval := time.Since(midnight) % job.GetRotation()

	gotlock := make(chan bool)

	go func () {
		job.Guard.Lock()
		log.Debug("Received lock for job", job.Msg.GetName())
		gotlock <- true
	}()

	select {
	case <- gotlock:
		go func() {
			log.Debug("Waiting to start", job.Msg.GetName())
			start := time.Now()
			defer log.Info("Finished job", job.Msg.GetName(), "took", time.Since(start))
			defer job.Guard.Unlock()
			timer := time.After(interval)
			for {
				select {
				case <- timer:
					log.Debug("Started job", job.Msg.GetName())
					// do job
				case <- *abort:
					log.Debug("Aborted job", job.Msg.GetName())
					// if, before we're done, the app exits, then GTFO.
					return
				}
			}
		}()
	case <- time.After(time.Second * 30):
		log.Debug("Couldn't get lock for job", job.Msg.GetName())
		return
	}
}

func (job Job) GetRotation() (rotate time.Duration) {
	freq := job.Msg.GetFrequency()
	if freq.GetHour() >= 0 {
		rotate += time.Hour * time.Duration(freq.GetHour())
	}
	if freq.GetMinute() >= 0 {
		rotate += time.Minute * time.Duration(freq.GetMinute())
	}
	if freq.GetSecond() >= 0 {
		rotate += time.Second * time.Duration(freq.GetSecond())
	}
	log.Debug("time.Duration for", job.Msg.GetName(), "is", rotate,
	"for input",
	freq.GetHour(), "h",
	freq.GetMinute(), "m",
	freq.GetSecond(), "s")
	return
}

func (job Job) HappensOn(t *time.Time) (bool) {
	freq := job.Msg.GetFrequency()

	return int8(freq.GetWeekday()) == int8(t.Weekday()) ||
	freq.GetWeekday() == -1 &&
	freq.GetMonth() == int32(t.Month()) &&
	freq.GetDay() == int32(t.Day())
}
