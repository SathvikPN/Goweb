package services

import "time"

func GetCurrentTime() time.Time {
	now := time.Now().UTC()
	kolkataLoc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}
	kolkataTime := now.In(kolkataLoc)
	return kolkataTime
}
