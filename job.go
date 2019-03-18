package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jasonlvhit/gocron"
)

// MonitorJob type
type MonitorJob struct {
	jobName string
	tenant  Tenant
	lastRun time.Time
	sender  *EmailSender
}

// NewJob and Start
func NewJob(t Tenant, e *EmailSender) *MonitorJob {
	j := &MonitorJob{
		jobName: fmt.Sprintf("CPI error message check for %s", t.Host),
		tenant:  t,
		lastRun: time.Now(),
		sender:  e,
	}
	if checkPassed, msg := CheckAPIAvailable(t); checkPassed {
		gocron.Every(uint64(t.Interval)).Seconds().Do(j.checkError)
		log.Printf("setup job for %s tenant", t.Host)
	} else {
		log.Printf(msg)
	}

	return j
}

func (j *MonitorJob) checkError() {
	now := time.Now()
	if msg, err := GetFailedInformationFor(j.tenant, j.lastRun); err == nil {
		if errCount, _ := strconv.ParseInt(*(msg.D.Count), 10, 64); errCount > 0 {

			notification := NotificationModel{
				Tenant:  j.tenant,
				LastRun: formatTime(j.lastRun),
				Now:     formatTime(now),
			}

			notification.Artifacts = GroupResultToArtifacts(msg.D.Results)

			for _, contact := range j.tenant.Contact {
				notification.ContactName = contact.Name
				j.sender.SendEmail(EmailPayload{
					To:      []string{contact.Email},
					Content: FormatTemplate(notification),
				})
			}
		}
	} else {
		log.Println(err)
		log.Printf("Get infromation from %s failed, please check the tenant status", j.tenant.Host)
	}

}

// StartAllJobs func
func StartAllJobs(config Config) {

	sender := NewSender(config.SMTP)

	for _, t := range config.Tenants {
		NewJob(t, sender)
	}

	log.Printf("starting jobs")

	<-gocron.Start()

}
