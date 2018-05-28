package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/srishanbhattarai/nepcal/dateconv"
)

// Cheap testing.
var writer io.Writer = os.Stdout

// Flag list
var (
	dateFlag = flag.Bool("d", false, "Print only the date")
)

func init() {
	flag.Parse()
}

func main() {
	render(*dateFlag)
}

// Render decides what to show based on the flags.
func render(dateFlag bool) {
	if dateFlag {
		showDate(writer, time.Now())
	} else {
		cal := newCalendar()
		cal.Render(writer, time.Now())
	}
}

// showDate prints the current B.S. date
func showDate(w io.Writer, t time.Time) {
	yy, mm, dd := t.Date()

	bsyy, bsmm, bsdd := dateconv.ToBS(toTime(yy, mm, dd)).Date()

	fmt.Fprintf(w, "%s %d, %d\n", dateconv.BSMonths[int(bsmm)], bsdd, bsyy)
}

// toTime creates a new time.Time with the basic yy/mm/dd parameters.
func toTime(yy int, mm time.Month, dd int) time.Time {
	return time.Date(yy, mm, dd, 0, 0, 0, 0, time.UTC)
}
