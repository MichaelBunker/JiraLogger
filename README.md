# Jira TimeTrap Logger

A Golang script to log hours automatically to Jira via the ticket API.

### Prerequisites

TimeTrap installed and used locally via the standard sqlite db.
Timesheet labled 'work'
Tickets in the format of 'ticket-###/Comments on work'
Environment variable pointing to the db timetrap is using.
```
export TIMETRAP_DATABASE=/Users/user/.timetrap.db
```

```
$ t sheet work
Switching to sheet "work"
$ t in NCP-31/Started work on API integration
$ t out
Checked out of sheet "work"
$ t t
Timesheet: work
    Day                Start      End        Duration   Notes
    Mon Feb 05, 2018   08:15:00 - 09:27:25   1:12:25    NCP-31/Started work on API integration

```

## Getting Started

 * Clone repository to your computer.
 * Make script executable.
 * Execute script with desired flags.
 * Enter credentials when prompted for actual posting (non-dry run)

```
$ ./JiraLogger --dryRun --round
  +-----------+-----------------------------+---------+
  |  TICKET   |         TIME SPENT          | STATUS  |
  +-----------+-----------------------------+---------+
  | NCP-30    | 1 hour 15 minutes 0 second  | Dry Run |
  | LSM-126   | 45 minutes 0 second         | Dry Run |
  | NCP-31    | 1 hour 15 minutes 0 second  | Dry Run |
  | PROJ-1642 | 30 minutes 0 second         | Dry Run |
  | LSM-126   | 45 minutes 0 second         | Dry Run |
  | NCP-31    | 1 hour 30 minutes 0 second  | Dry Run |
  | NCP-31    | 3 hours 0 minute 0 second   | Dry Run |
  +-----------+-----------------------------+---------+

$ ./JiraLogger --dryRun
  +-----------+--------------------------------+---------+
  |  TICKET   |           TIME SPENT           | STATUS  |
  +-----------+--------------------------------+---------+
  | NCP-30    | 1 hour 12 minutes 25 seconds   | Dry Run |
  | LSM-126   | 34 minutes 58 seconds          | Dry Run |
  | NCP-31    | 1 hour 11 minutes 50 seconds   | Dry Run |
  | PROJ-1642 | 15 minutes 0 second            | Dry Run |
  | LSM-126   | 30 minutes 3 seconds           | Dry Run |
  | NCP-31    | 1 hour 28 minutes 34 seconds   | Dry Run |
  | NCP-31    | 2 hours 53 minutes 24 seconds  | Dry Run |
  +-----------+--------------------------------+---------+

```


## Uses the following stuff.

* [Golang](https://golang.org/)
* [TimeTrap](https://github.com/samg/timetrap) - Time tracker
* [Jira](https://www.atlassian.com/software/jira) - Jira

