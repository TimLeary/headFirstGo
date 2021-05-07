package main

import "fmt"

type Liters float64
type Gallons float64
type Milliliters float64

func (l *Liters) LitersToGallons() Gallons {
	return Gallons(*l * 0.264)
}

func MillilitersToGallons(m Milliliters) Gallons {
	return Gallons(m * 0.000264)
}

func (g *Gallons) GallonsToLiters() Liters {
	return Liters(*g * 3.785)
}

func (g *Gallons) GallonsToMilliliters() Milliliters {
	return Milliliters(*g * 3785.41)
}

func main()  {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	additionLiter := Liters(40.0)
	carFuel += additionLiter.LitersToGallons()
	additionalGallon := Gallons(30.0)
	busFuel += additionalGallon.GallonsToLiters()

	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
	fmt.Printf("Car fuel in milliliters: %0.1f \n", carFuel.GallonsToMilliliters())
}
