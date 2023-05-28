package date_db

import "time"

const (
	dateLayout = "2006-01-02T15:04:05Z"
	databaseDateLayout = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(dateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(databaseDateLayout)
}