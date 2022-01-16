package main

import (
	helheim_go "github.com/bogdanfinn/helheim-go"
	"log"
	"net/http"
)

const YourApiKey = "INSERT_HERE"

func main() {
	helheimClient, err := helheim_go.NewClient(YourApiKey, false, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("helheim client initiated")

	options := helheim_go.CreateSessionOptions{
		Browser: helheim_go.BrowserOptions{
			Browser:  "chrome",
			Mobile:   false,
			Platform: "windows",
		},
		Captcha: helheim_go.CaptchaOptions{
			Provider: "vanaheim",
		},
	}
	session, err := helheimClient.NewSession(options)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("session response:")
	log.Println(session)

	balance, err := helheimClient.GetBalance()

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("balance response:")
	log.Println(balance)

	reqOpts := helheim_go.RequestOptions{
		Method:  http.MethodGet,
		Url:     "https://www.genx.co.nz/iuam/",
		Options: make(map[string]string),
	}
	resp, err := session.Request(reqOpts)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("request response:")
	log.Println(resp)

	err = session.Delete()

	if err != nil {
		log.Println(err)
		return
	}
}
