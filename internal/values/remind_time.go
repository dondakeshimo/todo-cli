package values

import (
	"fmt"
	"time"
)

type RemindTime string

const (
	layoutMin = "2006/1/2 15:04"
	layoutDay = "2006/1/2"
)

func NewRemindTime(str string) (RemindTime, error) {
	tt, err := parseStringToTime(str)
	if err != nil {
		return "", err
	}

	return RemindTime(tt.Format(layoutMin)), nil
}

func (rt RemindTime) AddTime(duration TimeDuration) (RemindTime, error) {
	tt, err := parseStringToTime(string(rt))
	if err != nil {
		return "", err
	}

	return RemindTime(tt.Add(time.Duration(duration)).Format(layoutMin)), nil
}

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

