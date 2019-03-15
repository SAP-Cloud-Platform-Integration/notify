package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jasonlvhit/gocron"
)

type MonitorJob struct {
	jobName string
	tenant  Tenant
	lastRun time.Time
	sender  *EmailSender
}

// NewJob and Start
func NewJob(t Tenant, e *EmailSender) *MonitorJob {
	s := gocron.NewScheduler()
	j := &MonitorJob{
		jobName: fmt.Sprintf("CPI error message check for %s", t.Host),
		tenant:  t,
		lastRun: time.Now(),
		sender:  e,
	}
	s.Every(uint64(t.Interval)).Seconds().Do(j.checkError)
	<-s.Start()
	logrus.Infof("job for %s tenant started", t.Host)
	return j
}

func (j *MonitorJob) checkError() {
	now := time.Now()
	errs := GetFailedInformationFor(j.tenant, j.lastRun)

	if errCount, _ := strconv.ParseInt(*(errs.D.Count), 10, 64); errCount > 0 {

		notification := NotificationModel{
			Tenant:  j.tenant,
			LastRun: formatTime(j.lastRun),
			Now:     formatTime(now),
		}

		notification.Artifacts = GroupResultToArtifacts(errs.D.Results)

		for _, contact := range j.tenant.Contact {
			notification.ContactName = contact.Name
			j.sender.SendEmail(EmailPayload{
				To:      []string{contact.Email},
				Content: FormatTemplate(notification),
			})
		}
	}

}

func StartAllJobs(config Config) {

	sender := NewSender(config.SMTP)

	for _, t := range config.Tenants {
		NewJob(t, sender)
	}

}
