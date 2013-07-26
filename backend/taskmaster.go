package backend

import (
	"sync"
	"time"
)

type Job struct {
	Guard sync.Mutex
	Msg   JobMsg
}

func (job *Job) Run(abort *chan bool) {
	t := time.Now()
	if !job.HappensOn(&t) {
		log.Debug("Job %s doesn't happen today. Bailing out.", job.Msg.GetName())
		return
	}

	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	interval := time.Since(midnight) % job.GetRotation()

	gotlock := make(chan bool)

	go func() {
		job.Guard.Lock()
		log.Debug("Received lock for job %s", job.Msg.GetName())
		gotlock <- true
	}()

	select {
	case <-gotlock:
		go func() {
			log.Debug("Waiting to start %s", job.Msg.GetName())
			start := time.Now()
			defer log.Info("Finished job %s took %s", job.Msg.GetName(),
				time.Since(start))
			defer job.Guard.Unlock()
			timer := time.After(interval)
			for {
				select {
				case <-timer:
					log.Debug("Started job %s", job.Msg.GetName())
					// do job
				case <-*abort:
					log.Debug("Aborted job %s", job.Msg.GetName())
					// if, before we're done, the app exits, then GTFO.
					return
				}
			}
		}()
	case <-time.After(time.Second * 30):
		log.Debug("Couldn't get lock for job %s", job.Msg.GetName())
		return
	}
}

func (job *Job) GetRotation() (rotate time.Duration) {
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
	log.Debug("time.Duration for %s is %s for input %d:%d:%d", job.Msg.GetName(), rotate,
		freq.GetHour(), freq.GetMinute(), freq.GetSecond())
	return
}

func (job *Job) HappensOn(t *time.Time) bool {
	freq := job.Msg.GetFrequency()

	return int8(freq.GetWeekday()) == int8(t.Weekday()) ||
		freq.GetWeekday() == -1 &&
			freq.GetMonth() == int32(t.Month()) &&
			freq.GetDay() == int32(t.Day())
}

func NewInterval(interval string) (*FrequencyMsg, error) {
	fm := &FrequencyMsg{}
	var min int32 = 2
	fm.Minute = &min
	return fm, nil
}

func NewJob(executable, args, interval, name string) (Job, error) {
	job := Job{}
	job.Guard = sync.Mutex{}
	job.Msg = JobMsg{
		Name:     &name,
		ExecPath: &executable,
		Args:     &args,
	}

	int_msg, interval_error := NewInterval(interval)
	if interval_error != nil {
		return Job{}, interval_error
	}
	job.Msg.Frequency = int_msg

	return job, nil
}
