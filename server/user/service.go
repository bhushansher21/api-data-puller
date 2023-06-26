package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"main.go/payload"
)

const jsonExtension = ".json"

var (
	jsonMarshal = json.Marshal
	osRename    = os.Rename
)

// saveUserDetails function save third party api response in json file
func saveUserDetails(wg *sync.WaitGroup, mu *sync.Mutex, config payload.Config, userList payload.UserList) error {

	content, err := jsonMarshal(userList)
	if err != nil {
		return fmt.Errorf("Error while marshalling user details %v", err.Error())
	}

	mu.Lock()
	err = validateFile(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Writing output in json file.
	err = ioutil.WriteFile(config.OutputFileCreationPath, content, 0644)
	if err != nil {
		return fmt.Errorf("Error while saving user details in json file %v", err.Error())
	}
	mu.Unlock()
	wg.Done()
	fmt.Println("user details saved successfully")
	return nil
}

// validateFile function is validating file size and renaming file name.
func validateFile(config payload.Config) error {
	var size int64
	//Checking file size.
	fi, err := os.Stat(config.OutputFileCreationPath)
	if !os.IsNotExist(err) {
		size = fi.Size()
	}

	//Avoiding the issue of file oversizing
	//Checking if file already present then rename it for creating new file.
	if size > 0 {
		newNameOfFile := strings.Replace(config.OutputFileCreationPath, jsonExtension, time.Now().Format(config.FileNameFormat), 1)
		err = osRename(config.OutputFileCreationPath, newNameOfFile+jsonExtension)
		if err != nil {
			return fmt.Errorf("Error while renaming old json file %v", err.Error())
		}
	}
	return nil
}
