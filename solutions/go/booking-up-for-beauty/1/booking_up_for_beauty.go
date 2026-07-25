package booking

import (
    "fmt"
    "time"
    "strings"
)

func ParseDate(date string, formatString string) time.Time {
    fmt.Printf("date argument %s\n", date)
	const longForm = "1/2/2006 15:04:05"
    fmt.Printf("formatString %s\n",formatString)
	t, err := time.Parse(formatString, date)
    if nil != err {
        fmt.Println(err)
    }
    fmt.Println(t)
    return t
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    return ParseDate(date, "1/2/2006 15:04:05")
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    return ParseDate(date, "January 2, 2006 15:04:05").Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    t := ParseDate(date, "Monday, January 2, 2006 15:04:05")
    hour := t.Hour()
    return hour >= 12 && hour <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    t := ParseDate(date, "1/2/2006 15:04:05")
	formattedTime := t.Format("Monday, January 2, 2006, at 15:04.")
    
    builder := strings.Builder{}

    builder.WriteString("You have an appointment on ")
    builder.WriteString(formattedTime)
    return builder.String()
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    t := ParseDate(time.Now().Format("2006")+"-09-15", "2006-01-02")
    return t
}
