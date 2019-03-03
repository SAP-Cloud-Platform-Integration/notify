package main

import "testing"

func TestEmailSender_SendEmail(t *testing.T) {
	type args struct {
		payload EmailPayload
	}
	tests := []struct {
		name string
		s    *EmailSender
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SendEmail(tt.args.payload)
		})
	}
}
