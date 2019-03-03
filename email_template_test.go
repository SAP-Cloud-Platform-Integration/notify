package main

import (
	"testing"
)

func TestFormatTemplate(t *testing.T) {
	type args struct {
		data NotificationModel
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"template format",
			args{
				NotificationModel{
					Tenant{
						60,
						"sample",
						"mock-tmn.hci.cn1.hana.ondemand.com",
						"username",
						"password",
						[]Contact{{
							"Theo Sun",
							"theo.sun@outlook.com",
						}},
					},
					"Theo Sun",
					"2019-03-01",
					"2019-03-03",
					[]Artifact{
						Artifact{
							ArtifactName: "Test Iflow",
							Errors:       (*ParseMessageProcessingLogFromFile("./email_template_test.json")).D.Results,
						},
					},
				},
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FormatTemplate(tt.args.data)
		})
	}
}
