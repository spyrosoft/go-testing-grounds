package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	mailchimp "github.com/spyrosoft/mailchimp-go"
	"github.com/spyrosoft/mailchimp-go/lists/members"
)

type Settings struct {
	APIKey     string `json:"key"`
	Subscriber string `json:"subscriber"`
	ListId     string `json:"list-id"`
}

func main() {
	settings, err := loadSettings()
	panicOnErr(err)

	err = mailchimp.SetKey(settings.APIKey)
	panicOnErr(err)

	params := &members.NewParams{
		EmailAddress: settings.Subscriber,
		Status:       members.StatusSubscribed,
	}

	member, err := members.New(settings.ListId, params)
	panicOnErr(err)

	fmt.Printf("%+v\n", member)
}

func loadSettings() (settings Settings, err error) {
	rawSettings, err := ioutil.ReadFile("private/settings.json")
	panicOnErr(err)
	err = json.Unmarshal(rawSettings, &settings)
	panicOnErr(err)
	return
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
