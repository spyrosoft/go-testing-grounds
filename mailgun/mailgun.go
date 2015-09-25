package main

import (
	"fmt"
	"github.com/mailgun/mailgun-go"
	"os"
	"io/ioutil"
	"encoding/json"
	"strings"
//	"encoding/csv"
)

type Config struct {
	Domain string `json:"domain"`
	PrivateKey string `json:"private-key"`
	PublicKey string `json:"public-key"`
}

func initializeConfig() Config {
	newConfig := Config{}
	var receivedSettings []byte
	receivedSettings, error := ioutil.ReadAll( os.Stdin )
	exitOnError( error )
	error = json.Unmarshal( receivedSettings, &newConfig )
	exitOnError( error )
	return newConfig
}

func sendEmail() {
	config := initializeConfig()
	mailgunInstance := mailgun.NewMailgun( config.Domain, config.PrivateKey, config.PublicKey )
	messageContentsBytes, error := ioutil.ReadFile( "message.email" )
	exitOnError( error )
	messageContents := string( messageContentsBytes )
	messageContents = strings.Replace( messageContents, "RECIPIENT_NAME", "Batman", 1 )
	newMessage := mailgunInstance.NewMessage( "matthew@craftedinsantacruz.com", "Testing Mailgun", messageContents, "matthew@eventsantacruz.com" )
	_, _, error = mailgunInstance.Send( newMessage )
	exitOnError( error )
}

func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}

func main() {
	sendEmail(  )
	fmt.Println( "Success!!" )
}