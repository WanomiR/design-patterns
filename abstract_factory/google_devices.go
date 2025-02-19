package main

import (
	"fmt"
	"time"
)

type Pixel struct {
	model        string
	cameraPixels int
	batteryLevel int
}

func (p *Pixel) MakeCall() {
	fmt.Println("Making a call from a ", p.model)
}

func (p *Pixel) SendMessage() {
	fmt.Println("Sending a message from a ", p.model)
}

func (p *Pixel) TakePicture() {
	fmt.Println("Taking a picture with a ", p.model)
}

func (p *Pixel) GetBatteryLevel() int {
	return p.batteryLevel
}

type PixelWatch struct {
	model        string
	batteryLevel int
}

func (pw *PixelWatch) ShowTime() time.Time {
	fmt.Println("Showing time on a ", pw.model)
	return time.Now()
}

func (pw *PixelWatch) ShowNotifications() []string {
	fmt.Println("Showing notifications on a ", pw.model)
	return []string{"new message", "missed call", "urgent alert"}
}

func (pw *PixelWatch) GetBatteryLevel() int {
	return pw.batteryLevel
}
