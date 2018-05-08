package main

import (
	"os"

	"github.com/JustAnotherOrganization/l5424"
	"github.com/JustAnotherOrganization/l5424/min5424"
)

var log l5424.MinLog

func init() {
	var err error
	log, err = min5424.New("nil", l5424.EmergencyLvl.String())
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println(os.Stdout, l5424.NoticeLvl, "This is a simple implementation of the min5424.Logger")
	log.Printf(os.Stdout, l5424.AlertLvl, "It allows for simple leveled logging using an RFC5424 compliant interface\n")
}
