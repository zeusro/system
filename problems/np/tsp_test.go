package np

import (
	"fmt"
	"testing"
)

// ok  	github.com/zeusro/system/problems/np	0.322s
// ok  	github.com/zeusro/system/problems/np	0.810s 不存距离可以快点
func TestTravel(t *testing.T) {
	s := NewSalesman(usCities)
	current := RandomUSCity()
	plans := s.Travel(current, s.Plan)
	if !s.IsSolvable(usCities) {
		t.Fatal("旅行计划不可行")
	}
	for i, plan := range plans {
		fmt.Printf("%v:%+vD\n", i, plan)
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
		s := NewSalesman(usCities)
		current := RandomUSCity()
		_ = s.Travel(current, s.Plan)
	}
}
