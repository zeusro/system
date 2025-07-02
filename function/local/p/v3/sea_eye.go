package v3

import (
	"fmt"
	"sync"
)

// NewSwallowGarden 新燕花园
type NewSwallowGarden struct {
	Limit    int64 //时间限制，这里简化为秒单位 3600s
	Distance int64 //环形跑道长度，这里限制为m单位
	V        int64 //速度单位
}

// OnlyLoveYou 从前一辆自行车很慢，后座只能载一个妹子
func OnlyLoveYou() {
	// Sherlock Holmes
	// Dr. John H. Watson
	var limit int64 = 3600                                               //s
	var distance int64 = 600                                             //m
	sherlock := NewSwallowGarden{Limit: limit, Distance: distance, V: 3} //我说：“遇到富婆不要慌，先加个微信。你要先锻炼👇🏻下面的技术，相信你自己，一定能把她折服。”
	watson := NewSwallowGarden{Limit: limit, Distance: distance, V: 2}   //他说：“你是我这辈子遇过最想殴打的人。”
	p := sherlock.P(watson)
	fmt.Println(p)
}

// LoveYouAll111 现在的大货车很强，一车能载很多人
// 这其实是一个行星绕日模型求不动点的核心算法
func LoveYouAll111() {
	var limit int64 = 36000  //s
	var distance int64 = 600 //m
	zeusro := NewSwallowGarden{Limit: limit, Distance: distance, V: 3}
	watson := NewSwallowGarden{Limit: limit, Distance: distance, V: 2}
	hera := NewSwallowGarden{Limit: limit, Distance: distance, V: 5}
	np := zeusro.NP([]NewSwallowGarden{watson, hera})
	fmt.Println(np)
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

// NP 我们的征途是星辰大海
func (sun NewSwallowGarden) NP(stars []NewSwallowGarden) []int64 {
	var wg sync.WaitGroup
	timingsMap := sync.Map{}
	result := make([]int64, 0)
	for _, star := range stars {
		wg.Add(1)
		go func(s NewSwallowGarden) {
			defer wg.Done()
			cycle := s.Distance / s.V
			// fixme: 算法的瓶颈在这个循环
			for time := cycle; time < s.Limit; time += cycle {
				existing, _ := timingsMap.LoadOrStore(time, int64(1))
				if v, ok := existing.(int64); ok {
					timingsMap.Store(time, v+1)
				}
			}
		}(star)
	}
	wg.Wait()
	timingsMap.Range(func(key, value any) bool {
		if t, ok := key.(int64); ok {
			if v, ojbk := value.(int64); ojbk && v == (int64(len(stars))+1) {
				result = append(result, t)
			}
		}
		return true
	})
	return result
}
