package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mailgun/mailgun-go"
)

type Config struct {
	Domain     string `json:"domain"`
	Sender     string `json:"sender"`
	Recipient  string `json:"recipient"`
	PrivateKey string `json:"private-key"`
}

var (
	config Config
)

func main() {
	initializeConfig()
	sendEmail()
	fmt.Println("Success!!")
}

func initializeConfig() {
	var configJson []byte
	configJson, err := ioutil.ReadAll(os.Stdin)
	exitOnErr(err)
	err = json.Unmarshal(configJson, &config)
	exitOnErr(err)
}

func sendEmail() {
	mg := mailgun.NewMailgun(config.Domain, config.PrivateKey, "")

	messageBytes, err := ioutil.ReadFile("message.email")
	exitOnErr(err)

	message := string(messageBytes)
	message = strings.Replace(message, "RECIPIENT_NAME", "Batman", 1)

	mgMessage := mg.NewMessage(config.Sender, "Testing Mailgun", message, config.Recipient)
	_, _, err = mg.Send(mgMessage)
	exitOnErr(err)
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
