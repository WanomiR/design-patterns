package main

import "fmt"

func main() {
	porsche, _ := newCar("Porsche 911")
	if err := porsche.Refuel(50); err != nil {
		fmt.Println(err)
	}
	if err := porsche.Drive(); err != nil {
		fmt.Println(err)
	}

	toyota, _ := newCar("Toyota LC")
	if err := toyota.Refuel(150); err != nil {
		fmt.Println(err)
	}
	if err := toyota.Drive(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n%+v", porsche, toyota)
}
