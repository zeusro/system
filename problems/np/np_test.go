package np

import (
	"fmt"
	"testing"
)

/*
最南端：佛罗里达州南端（Key West 附近）约 24.5°N
最北端：美加边界，如明尼苏达州、蒙大拿州一带约 49°N
最西端：加利福尼亚州的西海岸（如圣地亚哥）约 -124.8°W
最东端：缅因州的东部靠近加拿大边界约 -66.9°W
纬度 24.5°N 到 49.4°N
经度 -124.8°W 到 -66.9°W
*/
func TestUScity(t *testing.T) {
	for _, city := range usCities {
		if !IsInContinentalUS(city.Coordinates.Latitude, city.Coordinates.Longitude) {
			t.Fatal(city)
		}
	}
	fmt.Println(len(usCities))
}
