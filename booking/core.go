package booking

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/deeper-x/togodo/memo"
)

// SelectEvents read db content
func SelectEvents() (memo.Events, error) {
	evts := memo.Events{}
	jFile, err := os.Open("./assets/db.json")
	if err != nil {
		log.Println(err)
		return evts, err
	}

	defer jFile.Close()

	content, err := ioutil.ReadAll(jFile)
	if err != nil {
		log.Println(err)
		return evts, err
	}

	json.Unmarshal(content, &evts)

	return evts, nil
}
