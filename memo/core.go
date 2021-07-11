package memo

import (
	"log"

	"github.com/deeper-x/togodo/communicator"
	"github.com/deeper-x/togodo/mtime"
)

// Event is the object to be saved, to be remembered
type Event struct {
	InEvt int64  `json:"InEvt"`
	Rcpt  string `json:"Rcpt"`
	Msg   string `json:"Msg"`
}

// Events is the db
type Events struct {
	Data []Event `json:"events"`
}

// sendN func is where communication is sent
func sendN(e Event) {
	if mtime.EventIsNow(e.InEvt) {
		log.Println("event fired")
		communicator.SendEmail(e.Rcpt, e.Msg)
	}
}

// Handler return routine memo
func Handler(e Events) func() {
	return func() {
		for k := range e.Data {
			go sendN(e.Data[k])
		}
	}
}
