package main

import (
	"./database"
	"./jira"
	"./output"
	"encoding/json"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strconv"
)

func main() {

	roundFlag, dryRunFlag, helpFlag := getFlags()

	if *helpFlag {
		fmt.Println("This script is used for logging time to Jira tickets via TimeTrap time tracker. \r\n")
		fmt.Println("This script assumes you have a sheet in Timetrap labeled `work` and add entries in the following format - `ticket-#/comment`")
		fmt.Println("For example - `NCP-27/Work on changes to controller` \r\n")
		fmt.Println("The following flags are possible: \r\n")
		flag.PrintDefaults()
		return
	}

	test := database.GetRecordsArray(*roundFlag)
	data := [][]string{}

	username, password := "", ""

	if !*dryRunFlag {
		username, password = readCredentials()
	}
	for _, log := range test {
		ticket  := database.GetTicketNumber(log)
		requestData, _ := json.Marshal(log)

		if !*dryRunFlag {
			resp := jira.SendRequest(requestData, ticket, username, password)
			data = append(data, []string{ticket, output.GetDisplayTime(log.TimeSpentSeconds), strconv.Itoa(resp.StatusCode)})
		} else {
			data = append(data, []string{ticket, output.GetDisplayTime(log.TimeSpentSeconds), "Dry Run"})
		}
	}

	output.DisplayLogs(data)
}

//Get cli flags
func getFlags() (*bool, *bool, *bool) {
	roundFlag := flag.Bool("round", false, "Round entries up to 15 minutes intervals.")
	dryRunFlag := flag.Bool("dryRun", false, "Dry run of logger that outputs what will be logged.")
	helpFlag := flag.Bool("help", false, "Prints this message")
	flag.Parse()

	return roundFlag, dryRunFlag, helpFlag
}

// Get user credentials from cli prompt.
func readCredentials() (string, string) {

	fmt.Print("Username: ")
	username, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	CheckErr(err)

	fmt.Println()

	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	CheckErr(err)

	return string(username), string(password)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}