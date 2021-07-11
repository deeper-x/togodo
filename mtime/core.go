package mtime

import (
	"log"
	"strconv"
	"time"

	"github.com/deeper-x/togodo/settings"
)

// NowMinute returns now, rounded to minute in order to be used in a crontab service
// called with minute tolerance
func NowMinute() (int64, error) {
	tsNow := time.Now()
	tsNowRounded := tsNow.Format(settings.UTM)

	tsNowInt, err := strconv.Atoi(tsNowRounded)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return int64(tsNowInt), nil
}

// EventIsNow check if input event timestamp is scheduled now, rounded to minute
func EventIsNow(YYYYMMDDHHmm int64) bool {
	nm, _ := NowMinute()

	log.Println("comparing", nm, YYYYMMDDHHmm)

	return nm == YYYYMMDDHHmm
}
