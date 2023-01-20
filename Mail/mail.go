package Mail

import (
	"example/movieSuggestion/utils"
	"fmt"
	"strings"

	"gopkg.in/gomail.v2"
)

var mailKeys = strings.Split(utils.GoDotEnvVariable("MailKeys"), ",")
var index = 0

// This function returns mail body
func MailInfo(username string, receiver string) string {
	mailContent := fmt.Sprintf("Your Account is Successfully Created..<br><br> <b>Your account details are : </b> <br><br> Username : %s <br><br> Email : %s <br> ", username, receiver)
	return mailContent
}

// This function is used to send mail telling that account is created
func SendMail(username string, receiver string, mailInfo string) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "admin@gmail.com")
	msg.SetHeader("To", receiver)
	msg.SetHeader("Subject", fmt.Sprintf("Hello %s ..! Your account has been created ", username))
	msg.SetBody("text/html", mailInfo)
	// msg.Attach("/home/User/cat.jpg")
	if index >= len(mailKeys) {
		fmt.Printf("Cannot send mail..\nLimit Reached..\n")
		return
	}
	n := gomail.NewDialer("smtp.gmail.com", 587, "mailtempmail3@gmail.com", mailKeys[index])
	index++

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		panic(err)
	}
}
