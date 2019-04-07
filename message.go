package main

import (
	"github.com/getsentry/raven-go"
)

// GroupResultToArtifacts func
func GroupResultToArtifacts(messages []Result) []Artifact {
	artifacts := []Artifact{}
	errorGroups := make(map[string][]Result)

	for _, m := range messages {
		if g, ok := errorGroups[string(m.IntegrationArtifact.Name)]; ok {
			errorGroups[string(m.IntegrationArtifact.Name)] = append(g, m)
		} else {
			errorGroups[string(m.IntegrationArtifact.Name)] = []Result{m}
		}
	}

	for k, v := range errorGroups {
		artifacts = append(artifacts, Artifact{
			ArtifactName: k,
			Errors:       v,
		})
	}
	return artifacts
}

// CaptureMessages & send to sentry
func CaptureMessages(tenant Tenant, msgs []Result) {

	for _, m := range msgs {
		errMsg := GetErrorLogFor(tenant, m)
		tenant := string(tenant.Host)
		artifact := string(m.IntegrationArtifact.Name)
		evDate := ParseODataTimeStamp(*m.LogEnd)
		raven.Capture(
			&raven.Packet{
				Message:    errMsg,
				Level:      raven.ERROR, // error log
				Timestamp:  raven.Timestamp(evDate),
				ServerName: tenant,
				Tags: []raven.Tag{
					raven.Tag{
						Key:   "Artifact",
						Value: artifact,
					},
					raven.Tag{
						Key:   "Tenant",
						Value: tenant,
					},
					raven.Tag{
						Key:   "Status",
						Value: string(*m.Status),
					},
				},
				Extra: raven.Extra{
					"Sender":        m.Sender,
					"Receiver":      m.Receiver,
					"TransactionId": *m.TransactionID,
					"WebLink":       *m.AlternateWebLink,
				},
			},
			map[string]string{},
		)
	}

}
