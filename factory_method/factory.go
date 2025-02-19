package main

import "errors"

type Porshce911 struct {
	car
}

func newPorsche911() automobile {
	return &Porshce911{
		car: car{
			model:   "Porsche 911",
			maxFuel: 50,
		},
	}
}

type ToyotaLC struct {
	car
}

func newToyotaLC() automobile {
	return &ToyotaLC{
		car: car{
			model:   "Toyota LC",
			maxFuel: 150,
		},
	}
}

func newCar(model string) (automobile, error) {
	switch model {
	case "Porsche 911":
		return newPorsche911(), nil
	case "Toyota LC":
		return newToyotaLC(), nil
	default:
		return nil, errors.New("unknown model")
	}
}
