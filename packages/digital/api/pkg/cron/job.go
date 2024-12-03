package cron

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type CronJob struct {
	ExpirateAt time.Time
	JobCallback func() error
	Cancellation func()
	RunEvery time.Duration
	Credential string
	Id int
}

type CronGlobalHandler struct {
	listing map[string]*CronJob
	Logger bool
}

var cron *CronGlobalHandler

func New(withLogger bool) *CronGlobalHandler {
	cron = &CronGlobalHandler{
		listing: make(map[string]*CronJob),
		Logger: withLogger,
	}

	return cron
}


func (cron *CronGlobalHandler) CreateJob(ctx context.Context, credential string, when time.Duration, expire time.Time, callback func() error) (*CronJob, error) {
	ctx, cancel := context.WithCancel(ctx)
	newJob := &CronJob{
		ExpirateAt:   expire,
		JobCallback:  callback,
		Cancellation: cancel,
		Credential:   credential,
		Id:           len(cron.listing),
		RunEvery: when,
		}

		if cron.listing[credential] != nil {
			return nil, errors.New("Cron job with this credential is already exists")
		}

		cron.listing[credential] = newJob

		go func(ctx context.Context, job *CronJob) {
			durationExpire := time.Until(job.ExpirateAt)
			durationJob := job.RunEvery
			ticker := time.NewTicker(durationJob)
			timeAfter := time.After(durationExpire)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("\nCRON JOB %s is cancelled\n", job.Credential)
		            delete(cron.listing, job.Credential)
		            return // Exit the goroutine
				case <- timeAfter:
					delete(cron.listing, job.Credential)
					fmt.Printf("\nCRON JOB %s is ended successfully", job.Credential)
					cancel()
					break
					case <- ticker.C:
						if err := callback(); err != nil {
							delete(cron.listing, job.Credential)
							cancel()
						}
						break
				}
			}
		}(ctx, newJob)


		return newJob, nil
}



func (cron *CronGlobalHandler) EndWithDelete(cred string) {
	cron.listing[cred].Cancellation()
	delete(cron.listing, cred)
}
