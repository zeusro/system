package v1

import (
	"fmt"
	"testing"
)

// ok  	github.com/zeusro/system/problems/np/model/v1	0.319s
func TestTravel(t *testing.T) {
	s := NewSalesman(USCities)
	current := RandomUSCity()
	// current := USCities[0] // 直接指定起点城市，避免随机性
	// _ = s.Travel(current)
	plans := s.Travel(current)
	if !s.IsSolvable(USCities) {
		t.Fatal("旅行计划不可行")
	}
	for i, plan := range plans {
		fmt.Printf("%v:%+v\n", i, plan)
	}
	fmt.Printf("跨越漫长的旅程（%v km），终于见到KURO\n", s.KURO)
}

// BenchmarkTravel-8   	   16104	     76409 ns/op	   23016 B/op	      17 allocs/op
// 表示测试运行时使用的 GOMAXPROCS 数（8 个逻辑 CPU）
// b.N 的值，即测试框架实际执行这个函数的次数
// 平均每次执行耗时（单位：纳秒/每次操作），即每次 Travel() 的耗时
// 每次操作平均分配的内存字节数（Bytes per operation）
// 每次操作平均分配了几次内存（Allocations per operation）
func BenchmarkTravel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewSalesman(USCities)
		current := RandomUSCity()
		_ = s.Travel(current)
	}
}

// // ok  	github.com/zeusro/system/problems/np/model/v1	0.154s
// func TestTravelN(t *testing.T) {
// 	s := NewSalesman(USCities)
// 	// current := ConvertCityFromV1(v1.RandomUSCity())
// 	current := USCities[0]
// 	// 直接指定起点城市，避免随机性
// 	s.TravelN(current.Name, len(s.TodoCity))
// 	if !s.IsSolvable(USCities) {
// 		t.Fatal("旅行计划不可行")
// 	}
// 	start := 0
// 	for i := len(s.Plan) - 1; i > 0; i-- {
// 		fmt.Printf("%v:%+v\n", start, s.Plan[i])
// 		start++
// 	}
// 	fmt.Printf("跨越漫长的旅程（%v km），终于见到KURO\n", s.KURO)
// }

// // BenchmarkTravelN-8   	   15686	     74104 ns/op	   14696 B/op	       9 allocs/op
// func BenchmarkTravelN(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		s := NewSalesman(USCities)
// 		current := RandomUSCity()
// 		s.TravelN(current.Name, len(s.TodoCity))
// 	}
// }
