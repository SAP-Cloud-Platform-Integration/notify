package main

import (
	"reflect"
	"testing"
	"time"
)

func Test_formatTime(t *testing.T) {
	type args struct {
		t time.Time
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"simple format",
			args{MustParseRFC3339("2012-05-29T09:13:28")},
			"2012-05-29T09:13:28",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatTime(tt.args.t); got != tt.want {
				t.Errorf("formatTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseODataTimeStamp(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			"time simple format",
			args{
				"/Date(1551592817883)/",
			},
			time.Unix(0, 1551592817883*1000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseODataTimeStamp(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseODataTimeStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
