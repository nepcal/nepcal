package main

import (
	"fmt"
	"io"
	"text/tabwriter"
	"time"

	"github.com/srishanbhattarai/nepcal/dateconv"
)

// calendar struct represents the state required to render the B.S. calendar using a tabwriter
// that writes out to an io.Writer.
type calendar struct {
	val int
	w   *tabwriter.Writer
}

// newCalendar returns a new instance of calendar with the initial value of 1 with the provided io.Writer.
func newCalendar(w io.Writer) *calendar {
	tabw := tabwriter.NewWriter(w, 0, 0, 1, ' ', 0)

	return &calendar{
		val: 1,
		w:   tabw,
	}
}

// Render prints the BS calendar for the given time.Time.
// For printing formatted/aligned output, we use a tabwriter from the
// standard library. It doesn't support ANSI escapes so we cant have
// color/other enhancements to the output.(https://github.com/srishanbhattarai/nepcal/issues/4)
func (c *calendar) Render(ad time.Time) {
	bs := dateconv.ToBS(ad)

	c.renderBSDateHeader(bs)
	c.renderStaticDaysHeader()
	c.renderFirstRow(ad, bs)
	c.renderCalWithoutFirstRow(bs)

	c.w.Flush()
}

// renderFirstRow renders the first row of the calendar. The reason this needs
// to be handled separately is because there is a offset in each month which
// determines which day the month starts from - we need to tab space the 'offset' number
// of days, then start printing from the day after the offset.
func (c *calendar) renderFirstRow(ad time.Time, bs dateconv.BSDate) {
	offset := bs.MonthStartsAtDay(ad)
	for i := 0; i < offset; i++ {
		fmt.Fprintf(c.w, "\t")
	}

	for i := 0; i < (7 - offset); i++ {
		fmt.Fprintf(c.w, "\t%d", c.val)
		c.next()
	}

	fmt.Fprint(c.w, "\n")
}

// renderCalWithoutFirstRow renders the rest of the calendar without the first row.
// renderFirstRow will handle that due to special circumstances. We basically loop over
// each row and print 7 numbers until we are at the end of the month.
func (c *calendar) renderCalWithoutFirstRow(bs dateconv.BSDate) {
	daysInMonth, _ := bs.DaysInMonth() // Invariant: the month value is always valid and 'ok' is never false.

	for c.val < daysInMonth {
		start := daysInMonth - c.val
		end := start + 7

		for i := start; i < end; i++ {
			if c.val > daysInMonth {
				break
			}

			fmt.Fprintf(c.w, "\t%d", c.val)
			c.next()
		}

		fmt.Fprint(c.w, "\n")
	}
}

// renderStaticDaysHeader prints the static list of days for the calendar
func (c *calendar) renderStaticDaysHeader() {
	for _, v := range []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"} {
		fmt.Fprintf(c.w, "%s\t", v)
	}

	fmt.Fprint(c.w, "\n")
}

// renderBSDateHeader prints the date corresponding to the time e. This will
// be the header of the calendar.
func (c *calendar) renderBSDateHeader(e dateconv.BSDate) {
	yy, mm, dd := e.Date()

	if month, ok := dateconv.GetBSMonthName(time.Month(mm)); ok {
		fmt.Fprintf(c.w, "\t\t%s %d, %d\n\t", month, dd, yy)
	}
}

// next increments the value counter.
func (c *calendar) next() {
	c.val++
}
