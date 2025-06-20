package v2

import (
	"fmt"
	"testing"

	v1 "github.com/zeusro/system/problems/np/model/v1"
)

// ok  	github.com/zeusro/system/problems/np/model/v2	0.514s
func TestTravelN(t *testing.T) {
	s := NewSalesman(v1.USCities)
	// current := ConvertCityFromV1(v1.RandomUSCity())
	current := v1.USCities[0]
	// 直接指定起点城市，避免随机性
	s.TravelN(current.Name, len(s.TodoCity))
	if !s.IsSolvable(v1.USCities) {
		t.Fatal("旅行计划不可行")
	}
	start := 0
	for i := len(s.Plan) - 1; i > 0; i-- {
		fmt.Printf("%v:%+v\n", start, s.Plan[i])
		start++
	}
	fmt.Printf("跨越漫长的旅程（%v km），终于见到KURO\n", s.KURO)
}

// BenchmarkTravelN-8   	   15686	     74104 ns/op	   14696 B/op	       9 allocs/op
func BenchmarkTravelN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewSalesman(v1.USCities)
		current := v1.RandomUSCity()
		s.TravelN(current.Name, len(s.TodoCity))
	}
}
