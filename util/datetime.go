package util

import (
	"strconv"
	"time"
)

func TimeFromID(id string) (time.Time, error) {
	const offset = 1420070400000

	n, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	timestamp := n>>22 + offset

	var (
		s  = int64(timestamp / 1000)
		ns = int64(timestamp % 1000 * 1000000)
	)

	return time.Unix(s, ns), nil
}

func Midnight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, t.Location())
}
