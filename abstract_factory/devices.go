package main

import (
	"time"
)

type smartphone interface {
	MakeCall()
	SendMessage()
	TakePicture()
	GetBatteryLevel() int
}

type smartwatch interface {
	ShowTime() time.Time
	ShowNotifications() []string
	GetBatteryLevel() int
}
