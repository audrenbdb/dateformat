package dateformat

import (
	"fmt"
	"time"
)

var days = [7]string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"}
var months = [12]string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"}

//PrettyFrench takes a date and returns a string as follow "le lundi 01 janvier à 10h00"
func PrettyFrench(date *time.Time) string {
	return fmt.Sprintf("le %s %s %s à %sh%s",
		days[int(date.Weekday())],
		fmt.Sprintf("%02d", date.Day()),
		months[int(date.Month())-1],
		fmt.Sprintf("%02d", date.Hour()),
		fmt.Sprintf("%02d", date.Minute()))
}

//PrettyFrenchLong takes a date and returns a string as follow "le lundi 01 janvier 2019"
func PrettyFrenchLong(date *time.Time) string {
	return fmt.Sprintf("le %s %s %s %d",
		days[int(date.Weekday())],
		fmt.Sprintf("%02d", date.Day()),
		months[int(date.Month())-1],
		date.Year())
}

//PrettyFrenchShort takes a date and returns a string as follow "le lundi 01 janvier"
func PrettyFrenchShort(date *time.Time) string {
	return fmt.Sprintf("le %s %s %s",
		days[int(date.Weekday())],
		fmt.Sprintf("%02d", date.Day()),
		months[int(date.Month())-1])
}

//French formats as follow " le 01/01/2019 à 09h00"
func French(date *time.Time) string {
	return fmt.Sprintf("le %02d/%02d/%d à %02dh00",
		date.Day(),
		date.Month(),
		date.Year(),
		date.Hour())
}

//SetTime sets a date hour and minute based on a mn integer.
//I.E mn 630 is 10:30.
func SetTime(date time.Time, mn int) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), mn/60, mn%60, 0, 0, time.UTC)
}

//ToString returns a date string based on a date and a mn integer.
//I.E mn 630 is 10:30:00.
func ToString(date time.Time, mn int) string {
	return time.Date(date.Year(), date.Month(), date.Day(), mn/60, mn%60, 0, 0, time.UTC).Format("2006-01-02 15:04:05")
}

//StringToPrettyFrench takes a date string and a mn int and return string format as follow "le lundi 01 janvier à 10h00"
func StringToPrettyFrench(dateString string, mn int) (string, error) {
	var prettyDate string
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return prettyDate, err
	}
	date = SetTime(date, mn)
	prettyDate = PrettyFrench(&date)
	return prettyDate, nil
}
