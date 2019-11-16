//This is just demo version of facebook -> google events sync app
package main

import (
	"flag"
	"fmt"
	"log"

	fb "github.com/huandu/facebook"
)

//Response is a response json structure from API call
type Response struct {
	Events Events `json:"events"`
	ID     string `json:"id"`
}

//Events represents a data structure of all page, group or user events
type Events struct {
	Data   []Datum `json:"data"`
	Paging Paging  `json:"paging"`
}

//Datum holds an information about specifical event
type Datum struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

//Paging ...
type Paging struct {
	Cursors Cursors `json:"cursors"`
}

//Cursors ...
type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

var groupID = flag.String("group", "2484779941591369", "copy/paste your fb group id")
var accessToken = flag.String("token", "", "copy/paste your fb access token")

func main() {
	flag.Parse()
	if *accessToken == "" {
		log.Fatalln("Access token should be specified !")
	}

	/* 1. Retrieving all events list from fb */
	res, err := fb.Get(fmt.Sprintf("/%s", *groupID), fb.Params{
		"fields":       "events{name,description}",
		"access_token": *accessToken,
	})
	if err != nil {
		log.Fatal(err)
	}

	event := Response{
		Events: Events{
			Data: make([]Datum, 3),
			Paging: Paging{
				Cursors: Cursors{},
			},
		},
		ID: "",
	}

	/* 2. Unmarshaling list of fb events to structures */
	err = res.Decode(&event)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range event.Events.Data {
		fmt.Printf("Name: %s\n", v.Name)
		fmt.Printf("Description: %s\n", v.Description)
		fmt.Printf("Id: %s\n", v.ID)

		fmt.Println()
	}

	/* 3. Uploading fb events data to Google calendar */
	//TO BE DONE
}
