package store

import (
	"time"
	"errors"
	"strconv"
	"fmt"
)


func getTime(month int64) (time.Time, error) {
	if month < 1 || month > 12 {
		return time.Time{}, errors.New("Not valid value")
	}

	var stringNumberMonth string

	if month < 10 {
		stringNumberMonth = "0" + strconv.FormatInt(month, 10)
	} else {
		stringNumberMonth = strconv.FormatInt(month, 10)
	}

	fmt.Println(stringNumberMonth)
	layout:= time.DateOnly
	timeValue := fmt.Sprintf("%d-%s-01", time.Now().Year(), stringNumberMonth)


	newTime, err := time.Parse(layout, timeValue)

	if err != nil {
		return	time.Time{}, fmt.Errorf("Error parse new Time %w", err)
	}


	return newTime, nil
}

func getTimeFromString(s string) (time.Time, error) {
	layout:= time.DateOnly
	newTime, err := time.Parse(layout, s)

	if err != nil {
		return	time.Time{}, fmt.Errorf("Error parse new Time %w", err)
	}

	return newTime, nil
}