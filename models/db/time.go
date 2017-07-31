package db

import "time"

func FormattedTime() string {
	return time.Now().UTC().Format("20170329T203752+0900")
}
