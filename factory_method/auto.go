package main

type automobile interface {
	Drive() error
	FuelLeft() float64
	Refuel(float64) error
}
