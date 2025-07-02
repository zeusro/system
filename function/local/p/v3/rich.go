package v3

import "fmt"

// OnlyLoveYou 从前一辆自行车很慢，后座只能载一个妹子
func OnlyLoveYou() {
	// Sherlock Holmes
	// Dr. John H. Watson
	var limit int64 = 3600   //s
	var distance int64 = 600 //m
	zeusro := NewSwallowGarden{Limit: limit, Distance: distance, V: 3}
	watson := NewSwallowGarden{Limit: limit, Distance: distance, V: 2}
	p := zeusro.P(watson)
	fmt.Println(p)
}

// LoveYouAll111 现在的大货车很强，一车能载很多人
func LoveYouAll111() {
	// TODO
}

type NewSwallowGarden struct {
	Limit    int64 //时间限制，这里简化为秒单位 3600s
	Distance int64 //环形跑道长度，这里限制为m单位
	V        int64 //速度单位
}

// NewSwallowGarden n维世界求不动点
// 返回N维世界时间的不动点，单位是秒
// 时间复杂度：O(V1 + V2)（V1/V2 是速度）
// 空间复杂度：O(V1)
func (sherlock NewSwallowGarden) P(hera NewSwallowGarden) []int64 {
	timings := make(map[int64]bool)
	// 简化问题，时间=位移/速度
	cycle := sherlock.Distance / sherlock.V
	var time int64 = cycle
	for time = cycle; time < sherlock.Limit; time += cycle {
		timings[time] = true
	}
	result := make([]int64, 0)
	cycle = hera.Distance / hera.V
	time = cycle
	for time = cycle; time < hera.Limit; time += cycle {
		if timings[time] {
			result = append(result, time)
		}
	}
	return result
}
