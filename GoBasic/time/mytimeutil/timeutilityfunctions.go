package mytimeutil

import (
	"fmt"
	"time"
)

// utility function to convert RFC 3339 timestamp to time.time
func ConvertRFCTimeStampToTime(timestamp string) (timeType time.Time, err error) {
	timeType, e := time.Parse(time.RFC3339, timestamp)

	if e == nil {
		fmt.Println("timestamp is in the correct RFC format")
	} else {
		fmt.Println("Timestamp is in the wrong format ", e)
	}

	return timeType, e
}

// utility function to convert time.time to RFC 3339 timestamp
func ConvertTimeToRFCTimeStamp(timeType time.Time) (timestamp string) {
	timestamp = timeType.Format(time.RFC3339)
	return timestamp
}

// utility function to convert time.Time to unix epoch timestamp Milliseconds
func ConvertTimeToEpoch(timeType time.Time) (epoch int64) {
	epoch = timeType.Unix() * 1000
	return epoch
}

// utility function to convert unix epoch timestamp in Millisecond to time.Time
func ConvertEpochToTime(epoch int64) (timeType time.Time) {
	timeType = time.Unix(0, epoch*int64(time.Millisecond))
	return timeType
}
