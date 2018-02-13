package output

import (
	"os"
	"math"
	"github.com/olekukonko/tablewriter"
	"strconv"
)

func DisplayLogs(logs [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Ticket", "Time Spent", "Status"})
	for _, v := range logs {
		table.Append(v)
	}
	table.Render()
}

func plural(count int, singular string) (result string) {
	if (count == 1) || (count == 0) {
	    result = strconv.Itoa(count) + " " + singular + " "
	} else {
	    result = strconv.Itoa(count) + " " + singular + "s "
	}
	return
}

func GetDisplayTime(input int) (result string) {

	hours   := math.Floor(float64(input) / 60 / 60)
	seconds := input % (60 * 60)
	minutes := math.Floor(float64(seconds) / 60)
	seconds = input % 60

	if hours > 0 {
	    result = plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else if minutes > 0 {
	    result = plural(int(minutes), "minute") + plural(int(seconds), "second")
	} else {
	    result = plural(int(seconds), "second")
	}

	return
}