package main

import "errors"

type car struct {
	model   string
	fuel    float64
	maxFuel float64
}

func (c *car) Drive() error {
	if c.fuel < 10 {
		return errors.New("not enough fuel")
	}

	println("Driving ", c.model)
	c.fuel -= 10
	return nil
}

func (c *car) FuelLeft() float64 {
	return c.fuel
}

func (c *car) Refuel(amount float64) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}

	if c.fuel+amount > c.maxFuel {
		return errors.New("overflow")
	}

	c.fuel += amount
	return nil
}
