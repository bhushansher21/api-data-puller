package user

import (
	"fmt"
	"sync"
	"testing"

	"main.go/payload"
)

func TestGetUserDetails(t *testing.T) {
	tst := tests{
		{
			name: "Test if success response ",
			config: payload.Config{
				OutputFileCreationPath: "D:/Project/testData/output.json",
				ThirdPartyAPIEndpoint:  "https://24pullrequests.com/users.json",
			},
			want: "nil",
		},
		// {
		// 	name: "Test if marshalling failed",
		// 	config: payload.Config{
		// 		OutputFileCreationPath: "D:/Project/testData/output.json",
		// 		ThirdPartyAPIEndpoint:  "https://24pullrequests.com/users.json",
		// 	},
		// 	want: "Error while marshalling user details",
		// },
		// {
		// 	name: "Test if writing in file failed",
		// 	config: payload.Config{
		// 		OutputFileCreationPath: "",
		// 	},
		// 	want: "Error while saving user details in json file",
		// },
	}

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for _, tt := range tst {
		wg.Add(1)

		GetUserDetails(wg, mu, tt.config)
		// if gotErr != nil && !strings.Contains(gotErr.Error(), tt.want) {
		// 	t.Errorf("Expected %s but got %s", tt.want, gotErr.Error())
		// }
		fmt.Println("counter", tt)
	}
}
