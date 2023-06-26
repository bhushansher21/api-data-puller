package main

import (
	"strings"
	"testing"
)

type tests []struct {
	name string
	arg  string
	want string
}

func TestLoadConfig(t *testing.T) {
	tst := tests{
		{
			name: "Test if Read file fails",
			arg:  "input",
			want: "Error while reading config file open input",
		},
		{
			name: "Test if success response",
			arg:  "../testData/config.json",
			want: "nil",
		},
		{
			name: "Test if JSON UnMarshalling fails",
			arg:  "../testData/invalidConfig.json",
			want: "Error while unmarshling config file json",
		}}

	for _, tt := range tst {
		_, gotErr := loadConfig(tt.arg)
		if gotErr != nil && !strings.Contains(gotErr.Error(), tt.want) {
			t.Errorf("Expected %s but got %s", tt.want, gotErr.Error())
		}
	}
}
