package testing

import (
	"fmt"
	"log"
	"net/smtp"
)

const (
	sender   = "edmpulasky@mymail.com"
	password = "mypassword"
	hostname = "smtp.example.com"
	template = `Warning: you are using %d bytes of storage, %d%% of your quota.`
)

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", username, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smpt.sendMail(%s), err:%s", username, err)
	}

}

func bytesInUse(username string) int64 { return 0 }

func checkQuota(username string) {
	used := bytesInUse(username)
	const quota = 1e9
	percent := 100 * used / quota

	if percent < 90 {
		return
	}

	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}
