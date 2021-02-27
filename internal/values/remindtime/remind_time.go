package remindtime

import (
	"fmt"
	"time"
)

/*
RemindTime is a string than can be parsed time.Time in some layouts.
Allowed layouts are
    "2006/1/2 15:04"
    "2006/1/2"
*/
type RemindTime string

const (
	layoutMin = "2006/1/2 15:04"
	layoutDay = "2006/1/2"
)

// NewRemindTime is a function that construct RemindTime.
func NewRemindTime(str string) (RemindTime, error) {
	tt, err := parseStringToTime(str)
	if err != nil {
		return "", err
	}

	return RemindTime(tt.Format(layoutMin)), nil
}

// AddTime is a function that add (or substract) duration to RemindTime.
func (rt RemindTime) AddTime(rlt RelativeTime) (RemindTime, error) {
	if !rlt.IsTaskBase {
		return RemindTime(time.Now().Add(rlt.RelativeTime).Format(layoutMin)), nil
	}

	tt, err := parseStringToTime(string(rt))
	if err != nil {
		return "", err
	}

	return RemindTime(tt.Add(rlt.RelativeTime).Format(layoutMin)), nil
}

func (rt RemindTime) ToTime() time.Time {
	t, _ := time.ParseInLocation(layoutMin, string(rt), time.Local)
	return t
}

// parseStringToTime parse allowed layouts string to time.Time.
func parseStringToTime(str string) (time.Time, error) {
	tM, errM := time.ParseInLocation(layoutMin, str, time.Local)
	tD, errD := time.ParseInLocation(layoutDay, str, time.Local)

	if errM != nil && errD != nil {
		return time.Time{}, fmt.Errorf("invalid time layout: [minutes layout]: %s, [day layout]: %s", errM.Error(), errD.Error())
	} else if errM == nil && errD != nil {
		return tM, nil
	} else if errM != nil && errD == nil {
		return tD, nil
	}

	return time.Time{}, fmt.Errorf("Parse failed for some reason")
}
