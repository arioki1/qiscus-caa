package qiscusResponse

import (
	"strconv"
	"strings"
	"time"
)

type GetOfficeHours struct {
	Data struct {
		OnlineMessage          interface{} `json:"online_message"`
		OfflineMessage         interface{} `json:"offline_message"`
		Timezone               string      `json:"timezone"`
		SendOnlineIfResolved   bool        `json:"send_online_if_resolved"`
		SendOfflineEachMessage bool        `json:"send_offline_each_message"`
		OfficeHours            []struct {
			Id        int    `json:"id"`
			AppId     int    `json:"app_id"`
			Day       int    `json:"day"`
			Starttime string `json:"starttime"`
			Endtime   string `json:"endtime"`
			CreatedAt string `json:"created_at"`
			UpdatedAt string `json:"updated_at"`
		} `json:"office_hours"`
	} `json:"data"`
}

func (h *GetOfficeHours) IsOfficeHours() bool {
	result := false

	layout := "15:04"
	zone := h.parseTzFromOffset(h.Data.Timezone)
	currentTime := time.Now().In(zone)
	numberOfDay := int(currentTime.Weekday())
	currentClock, _ := time.Parse(layout, currentTime.Format("15:04"))

	for _, officeHour := range h.Data.OfficeHours {
		if officeHour.Day == numberOfDay {
			startTime, _ := time.Parse(layout, officeHour.Starttime)
			endTime, _ := time.Parse(layout, officeHour.Endtime)
			if currentClock.After(startTime) && currentClock.Before(endTime) {
				result = true
			}
		}
	}
	return result
}

func (h *GetOfficeHours) parseTzFromOffset(offsetStr string) *time.Location {
	offset := strings.Split(offsetStr, ":")

	hour := offset[0]
	hourInt, err := strconv.Atoi(hour)
	if err != nil {
		utc, _ := time.LoadLocation("UTC")
		return utc
	}

	minute := offset[1]
	minuteInt, err := strconv.Atoi(minute)
	if err != nil {
		utc, _ := time.LoadLocation("UTC")
		return utc
	}

	if strings.HasPrefix(offsetStr, "-") {
		return time.FixedZone(offsetStr, hourInt*60*60-60*minuteInt)
	} else {
		return time.FixedZone(offsetStr, hourInt*60*60+60*minuteInt)
	}
}
