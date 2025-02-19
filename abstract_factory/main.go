package main

import "fmt"

func main() {
	appleFactory := new(AppleFactory)
	googleFactory := new(GoogleFactory)

	iphone := appleFactory.CreateSmartphone()
	applewatch := appleFactory.CreateSmartwatch()

	pixel := googleFactory.CreateSmartphone()
	pixielWatch := googleFactory.CreateSmartwatch()

	fmt.Println("Apple iPhone:", iphone)
	fmt.Println("Apple Watch:", applewatch)
	fmt.Println("Google Pixel:", pixel)
	fmt.Println("Google Watch:", pixielWatch)
}
