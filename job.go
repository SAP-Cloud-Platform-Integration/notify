package main

import "time"

type MonitorJob struct {
	jobName string
	tenant  Tenant
	lastRun time.Time
}

func StartAllJobs(config Config) {

}
