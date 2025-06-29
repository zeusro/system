package n

import (
	"fmt"
	"testing"
)

func TestEarthLocation(t *testing.T) {
	el := EarthLocation{}
	el.Coordinates.Latitude = 40.7128
	el.Coordinates.Longitude = -74.0060

	e2 := EarthLocation{}
	e2.Coordinates.Latitude = 34.0522
	e2.Coordinates.Longitude = -118.2437
	distance := el.GetEarthDistance(&e2)
	fmt.Printf("distance struct: %#v\n", distance)
	fmt.Printf("el: %+v \n", distance.ValueFloat64)
	// fmt.Printf("distance: %+v \n", distance)
}
