package main

import (
	"fmt"
	"github.com/mailgun/mailgun-go"
	"os"
	"encoding/json"
//	"encoding/csv"
)

type Config struct {
	Domain string `json:"domain"`
	PrivateKey string `json:"private-key"`
	PublicKey string `json:"public-key"`
}

func initializeConfig() {
	var receivedSettings []byte
	receivedSettings, error := ioutil.ReadAll( os.Stdin )
	exitOnError( error )
	error = json.Unmarshal( received_settings, &settings )
	exitOnError( error )
	min_pixel_intensity = settings.MaxIterations
}

func exitOnError( error error ) {
	if error != nil {
		fmt.Println( error )
		os.Exit( 1 )
	}
}

func main() {
	Mailgun = mailgun.NewMailgun(conf.Mailgun.Domain, conf.Mailgun.PrivateKey, conf.Mailgun.PublicKey)
	fmt.Println( "Success!!" )
}