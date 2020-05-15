package main

import (
	"reflect"
	"testing"
)

func TestParseConfigFromPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			"basic test",
			args{"./sample_config.json"},
			&Config{
				SMTP: SMTP{
					Username: "username",
					Password: "password",
					Server:   "1.2.3.4",
					Port:     "465",
					From:     "Theo Sun <theo.sun@abc.com>",
				},
				Tenants: []Tenant{Tenant{
					60,
					"sample",
					"mock-tmn.hci.cn1.hana.ondemand.com",
					"username",
					"password",
					[]Contact{{
						"Theo Sun",
						"theo.sun@outlook.com",
					}},
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseConfigFromPath(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConfigFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
