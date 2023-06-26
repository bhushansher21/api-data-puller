package user

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"

	"main.go/payload"
)

type tests []struct {
	name   string
	config payload.Config
	arg    payload.UserList
	want   string
}

func fakeMarshal(v interface{}) ([]byte, error) {
	return []byte{}, errors.New("Error while marshalling user details")
}

func restoreMarshal(replace func(v interface{}) ([]byte, error)) {
	jsonMarshal = replace
}

func TestSaveUserDetails(t *testing.T) {
	tst := tests{
		{
			name: "Test if success response ",
			config: payload.Config{
				OutputFileCreationPath: "../testData/output.json",
			},
			arg: payload.UserList{
				{ID: 0001, Nickname: "Bhushan", GravatarID: "", GithubProfile: "https://github.com/Bhushan"},
			},
			want: "nil",
		},
		{
			name: "Test if marshalling failed",
			config: payload.Config{
				OutputFileCreationPath: "../testData/output.json",
			},
			arg: payload.UserList{
				{ID: 0001, Nickname: "Bhushan", GravatarID: "", GithubProfile: "https://github.com/Bhushan"},
			},
			want: "Error while marshalling user details",
		},
		{
			name: "Test if writing in file failed",
			config: payload.Config{
				OutputFileCreationPath: "",
			},
			arg: payload.UserList{
				{ID: 0001, Nickname: "Bhushan", GravatarID: "", GithubProfile: "https://github.com/Bhushan"},
			},
			want: "Error while saving user details in json file",
		},
	}

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for _, tt := range tst {
		var storedMarshal func(v any) ([]byte, error)
		if tt.name == "Test if marshalling failed" {
			storedMarshal = jsonMarshal
			jsonMarshal = fakeMarshal
		}
		wg.Add(1)

		gotErr := saveUserDetails(wg, mu, tt.config, tt.arg)
		if gotErr != nil && !strings.Contains(gotErr.Error(), tt.want) {
			t.Errorf("Expected %s but got %s", tt.want, gotErr.Error())
		}
		fmt.Println("counter", tt)

		if tt.name == "Test if marshalling failed" {
			restoreMarshal(storedMarshal)
		}
	}
}

func fakeOsRename(old, new string) error {
	return errors.New("Error while renaming old json file of link")
}

func restoreOsRename(replace func(old, new string) error) {
	osRename = replace
}

func TestValidateFile(t *testing.T) {
	tst := tests{
		{
			name: "Test if renaming of file fail",
			config: payload.Config{
				OutputFileCreationPath: "../testData/output.json",
			},
			want: "Error while renaming old json file",
		},
		{
			name: "Test if success response",
			config: payload.Config{
				OutputFileCreationPath: "../testData/dummyOutput.json",
			},
			want: "nil",
		},
	}

	for _, tt := range tst {
		if tt.name == "Test if renaming of file fail" {
			storedOsRename := osRename
			osRename = fakeOsRename
			defer restoreOsRename(storedOsRename)
		}
		gotErr := validateFile(tt.config)
		if gotErr != nil && !strings.Contains(gotErr.Error(), tt.want) {
			t.Errorf("Expected %s but got %s", tt.want, gotErr.Error())
		}
	}
}
