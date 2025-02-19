package main

import (
	"fmt"
	"time"
)

type IPhone struct {
	model        string
	cameraPixels int
	batteryLevel int
}

func (i *IPhone) MakeCall() {
	fmt.Println("Making a call from an ", i.model)
}

func (i *IPhone) SendMessage() {
	fmt.Println("Sending a message from an ", i.model)
}

func (i *IPhone) TakePicture() {
	fmt.Println("Taking a picture with an ", i.model)
}

func (i *IPhone) GetBatteryLevel() int {
	return i.batteryLevel
}

type AppleWatch struct {
	model        string
	batteryLevel int
}

func (aw *AppleWatch) ShowTime() time.Time {
	fmt.Println("Showing time on a ", aw.model)
	return time.Now()
}

func (aw *AppleWatch) ShowNotifications() []string {
	fmt.Println("Showing notifications on a ", aw.model)
	return []string{"new message", "missed call"}
}

func (aw *AppleWatch) GetBatteryLevel() int {
	return aw.batteryLevel
}
