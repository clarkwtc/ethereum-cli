package utils

import "time"

func ToUTCTime(timestamp uint64) time.Time {
    return time.Unix(int64(timestamp), 0).UTC()
}
