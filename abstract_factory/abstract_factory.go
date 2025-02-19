package main

type devicesFactory interface {
	CreateSmartphone() smartphone
	CreateSmartwatch() smartwatch
}

type AppleFactory struct{}

func (af *AppleFactory) CreateSmartphone() smartphone {
	return &IPhone{
		model:        "iPhone 12",
		cameraPixels: 12,
		batteryLevel: 100,
	}
}

func (af *AppleFactory) CreateSmartwatch() smartwatch {
	return &AppleWatch{
		model:        "Apple Watch Series 6",
		batteryLevel: 100,
	}
}

type GoogleFactory struct{}

func (gf *GoogleFactory) CreateSmartphone() smartphone {
	return &Pixel{
		model:        "Google Pixel 5",
		cameraPixels: 24,
		batteryLevel: 100,
	}
}

func (gf *GoogleFactory) CreateSmartwatch() smartwatch {
	return &PixelWatch{
		model:        "Pixel Watch",
		batteryLevel: 100,
	}
}
