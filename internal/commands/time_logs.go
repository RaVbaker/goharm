package commands

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
	"sort"
	"time"

	"github.com/ravbaker/goharm/internal/config"
	"github.com/ravbaker/goharm/internal/jsonapi/client"
	"github.com/ravbaker/goharm/internal/jsonapi/resources"
)

func TimeLogs(cfg *config.Config, rangeFilter string) {
	now := time.Now().AddDate(0, 0, -7)
	startsAtFilter := floorDate(&now, rangeFilter).Format(time.RFC3339)

	timeLogs := getTimeLogs(cfg.General.Authentication.UserId, startsAtFilter)
	if len(timeLogs) == 0 {
		fmt.Printf("No logs found for range %s - %s\n", rangeFilter, startsAtFilter)
		return
	}

	sort.Sort(byStartsAt(timeLogs))
	for i, timeLog := range timeLogs {
		line := timeLog.(*resources.TimeLog)
		println(i, line.StartsAt, line.EndsAt, line.Description)
	}
}

func getTimeLogs(currentUserId, startsAtFilter string) []interface{} {
	query := url.Values{}
	query.Add("filter[user-id]", currentUserId)
	query.Add("filter[starts-at]", startsAtFilter)

	results, err := client.RequestMany("/api/v1/time-logs?"+query.Encode(), nil, reflect.TypeOf(&resources.TimeLog{}))
	if err != nil {
		log.Fatalln(err.Error())
	}
	return results
}

type byStartsAt []interface{}

func (s byStartsAt) Len() int {
	return len(s)
}
func (s byStartsAt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byStartsAt) Less(i, j int) bool {
	log1 := s[i].(*resources.TimeLog).StartsAt
	log2 := s[j].(*resources.TimeLog).StartsAt
	return log1 < log2
}

// FloorDate returns the floor of a date based on the unit given
// "week" returns a date pointing to the first day of the week of the given timestamp t
// "month" returns a date pointing to the first day of the month of the given timestamp t

func floorDate(t *time.Time, unit string) time.Time {
	y, m, d := t.Date()
	isoy, isow := t.ISOWeek()
	zoneName, zoneOffset := t.Zone()

	var result time.Time
	switch unit {
	case "week":
		// Since a week can span across 2 different years, or
		// months, it is important to work with isoyear and
		// isoweek instead of just year
		result = firstDayOfISOWeek(isoy, isow, time.FixedZone(zoneName, zoneOffset))
	case "month":
		result = time.Date(y, m, 1, 0, 0, 0, 0, time.FixedZone(zoneName, zoneOffset))
	default:
		result = time.Date(y, m, d, 0, 0, 0, 0, time.FixedZone(zoneName, zoneOffset))
	}
	return result
}

func firstDayOfISOWeek(year int, week int, timezone *time.Location) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, timezone)
	isoYear, isoWeek := date.ISOWeek()
	for date.Weekday() != time.Monday { // iterate back to Monday
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoYear < year { // iterate forward to the first day of the first week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	for isoWeek < week { // iterate forward to the first day of the given week
		date = date.AddDate(0, 0, 1)
		isoYear, isoWeek = date.ISOWeek()
	}
	return date
}
