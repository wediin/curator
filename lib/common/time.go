package common

import (
	"strconv"
	"time"
)

func TransformRFC3339Time(rfc3339Timestamp string) string {
	t, err := time.Parse(time.RFC3339, rfc3339Timestamp)
	if err != nil {
		panic(err)
	}

	return strconv.FormatInt(t.Unix(), 10)
}
