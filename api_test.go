package main

import (
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestGetFailedInformationFor(t *testing.T) {
	type args struct {
		t    Tenant
		from time.Time
	}
	tests := []struct {
		name string
		args args
		want *MessageProcessingLog
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := GetFailedInformationFor(tt.args.t, tt.args.from)
			assert.Equal(t, len(got.D.Results), 100)
		})
	}
}
