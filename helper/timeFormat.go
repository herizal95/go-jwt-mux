package helper

import "time"

func TimeFormat(format string) string {

	t := time.Now()

	formatDate := t.Format("2006-01-02 15:04:05")

	return formatDate
}
