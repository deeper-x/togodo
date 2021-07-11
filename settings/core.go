package settings

import (
	"os"
)

// TickInterval event is fired
const TickInterval = 60

// UTM format - Unix TimeStamp Minutes expressed YYYYMMDDHHmm
const UTM = "200601021504"

// EmailPasswd SMTP password
var EmailPasswd = os.Getenv("EMAIL_PASSWD")

// SMTPHOst gmail server
var SMTPHOst = os.Getenv("EMAIL_HOST")

// SMTPURI gmail server:port
var SMTPURI = os.Getenv("EMAIL_URI")

// SMTPUser notification sender
var SMTPUser = os.Getenv("EMAIL_USERNAME")
