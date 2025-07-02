package v2

import (
	"fmt"
	"testing"
)

func TestLcmOfList(t *testing.T) {
	// 圆形跑道长度（单位：米）
	trackLength := 600

	// 妹子A，富婆B，少妇C，淑女D，皮神的速度（单位：米/秒）
	speeds := []int{1, 2, 3, 4, 5}

	// 每人跑完一圈的时间 = 赛道长度 / 速度
	times := make([]int, len(speeds))
	for i, v := range speeds {
		times[i] = trackLength / v
	}

	// 求这些时间的最小公倍数
	result := LcmOfList(times)

	fmt.Printf("所有人在 %d 秒后会再次在起点相遇。\n", result)
}
