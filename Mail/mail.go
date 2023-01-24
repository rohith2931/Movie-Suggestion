package Mail

import (
	"example/movieSuggestion/utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var AdminMail = "admin@gmail.com"

// This function returns mail body
func MailInfo(username string, receiver string) string {
	mailContent := fmt.Sprintf(`{
	    "personalizations": [
	        {
	            "to": [
	                {
	                    "email": `+`"%s"`+`
	                }
	            ],
	            "subject": "Hello `+`%s`+`"
	        }
	    ],
	    "from": {
	        "email": `+`"%s"`+`
	    },
	    "content": [
	        {
	            "type": "text/plain",
	            "value":`+` "Your Account is Successfully Created .\n\n Your Account details are :\n\n Username : `+fmt.Sprint(username)+` \n\n Email : `+fmt.Sprint(receiver)+`"
			}
	    ]
	}`, receiver, username, AdminMail)
	return mailContent
}

// This function is used to send mail telling that account is created
func SendMail(mail string) {
	url := utils.GoDotEnvVariable("RapidAPI_URL")

	payload := strings.NewReader(mail)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", utils.GoDotEnvVariable("X-RapidAPI-Key"))
	req.Header.Add("X-RapidAPI-Host", "rapidprod-sendgrid-v1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	utils.CheckError(err)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
