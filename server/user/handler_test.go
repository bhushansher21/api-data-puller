package user

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"main.go/payload"
)

func TestGetUserDetails(t *testing.T) {

	// generate a test server so we can capture and inspect the request
	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("body"))
	}))
	defer func() { testServer.Close() }()

	tst := tests{
		{
			name: "Test if success response ",
			config: payload.Config{
				OutputFileCreationPath: "../../testData/output.json",
				ThirdPartyAPIEndpoint:  "https://24pullrequests.com/users.json",
			},
			want: "nil",
		},
	}

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for _, tt := range tst {
		wg.Add(1)
		req, err := http.NewRequest(http.MethodGet, testServer.URL, nil)
		assert.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode, "status code should match the expected response")

		GetUserDetails(wg, mu, tt.config)
	}
}
