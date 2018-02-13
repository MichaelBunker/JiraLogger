package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

// Struct for Log sent to Jira for creating new worklogs.
type JiraLog struct {
	Comment          string `json:"comment"`
	TimeSpentSeconds int `json:"timeSpentSeconds"`
}

func GetRecordsArray(round bool) []JiraLog {
	rows := GetDailyJiraLogs()
	data := []JiraLog{}
	for rows.Next() {
		logRecord := GetlogRecord(rows, round)
		data = append(data, logRecord)
	}
	return data
}

// Get the JiraLog for a given sql.Row
func GetlogRecord(row *sql.Rows, round bool) JiraLog {
	log := JiraLog{}
	err := row.Scan(&log.Comment, &log.TimeSpentSeconds)
	CheckErr(err)
	if round {
		log.TimeSpentSeconds = RoundToQuarterHour(log.TimeSpentSeconds)
	}

	return log
}

func GetTicketNumber(log JiraLog) string {
	return strings.Split(log.Comment, "/")[0]
}

//Round seconds to nearest quarter hour.
func RoundToQuarterHour(duration int) int {
	return ((duration + 900) / 900 ) * 900;
}

func GetDailyJiraLogs() *sql.Rows {

	CheckEnvVars("TIMETRAP_DATABASE")
	db, err := sql.Open("sqlite3", os.Getenv("TIMETRAP_DATABASE"))
	CheckErr(err)

	queryString := GetQuery()
	rows, err 	:= db.Query(queryString)
	CheckErr(err)

	return rows
}

// Build query string for sqlLite3 db. Getting all the entries and their durations for today.
func GetQuery() string {
	current_time := time.Now().Local()
	query := "SELECT note, ((strftime('%s', end) - strftime('%s', start))) as duration FROM entries where start > date('%s') AND sheet = 'work' AND end NOT NULL"

	return fmt.Sprintf(query, "%s", "%s", fmt.Sprintf("%d-%02d-%02dT00:00:01-00:00", current_time.Year(), current_time.Month(), current_time.Day()))
}

func CheckEnvVars(key string) {
	_, defined := os.LookupEnv(key)
	if !defined {
		fmt.Println("Environment key not set, %s", key)
		os.Exit(1)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}