package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"main.go/payload"
)

// GetUserDetails function invoke third party api and laod user details.
func GetUserDetails(wg *sync.WaitGroup, mu *sync.Mutex, config payload.Config) {
	request, err := http.NewRequest(http.MethodGet, config.ThirdPartyAPIEndpoint, nil)
	if err != nil {
		fmt.Errorf("Error while invoking third party API %v", err.Error())
	}

	res, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Errorf("Error while invoking third party API %v", err.Error())
	}
	defer res.Body.Close()

	userList := payload.UserList{}
	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &userList)
	if err != nil {
		fmt.Errorf("Error while unmarshalling third party API %v", err.Error())
	}

	err = saveUserDetails(wg, mu, config, userList)
	if err != nil {
		log.Fatal(err.Error())
	}
}
