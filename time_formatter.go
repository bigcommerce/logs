package logs

import "time"

var DefaultTimeFormatter = &TimeFormatter{
	TimeFormat: "02 Jan 2006 15:04:05.000 MST",
	Location:   time.UTC,
}

type TimeFormatter struct {
	TimeFormat string
	Location   *time.Location
}

func NewTimeFormatter(format string, location *time.Location) *TimeFormatter {
	return &TimeFormatter{
		Location:   location,
		TimeFormat: format,
	}
}

func (t *TimeFormatter) Format(event *Event) []byte {
	var (
		format = t.TimeFormat
		loc    = t.Location
	)
	if format == "" {
		format = DefaultTimeFormatter.TimeFormat
	}
	if loc == nil {
		loc = DefaultTimeFormatter.Location
	}
	return []byte(event.Time.In(loc).Format(format))
}