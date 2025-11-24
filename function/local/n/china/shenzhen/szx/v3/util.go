package v3

import (
	"fmt"
	"time"
)

func printGradient(text string) {
	// 简单渐变：从紫到蓝到青
	colors := []string{
		"\033[38;2;180;0;255m",  // 紫
		"\033[38;2;100;80;255m", // 蓝紫
		"\033[38;2;80;120;255m", // 蓝
		"\033[38;2;40;200;255m", // 青蓝
		"\033[38;2;0;255;255m",  // 青
	}
	reset := "\033[0m"

	runes := []rune(text)
	step := len(colors)
	for i, r := range runes {
		c := colors[i%step]
		fmt.Printf("%s%c%s", c, r, reset)
	}
	fmt.Println()
}

// FindMinTime 返回切片 ts 中第一个非零 time.Time 的索引和最小时间。
// 如果全部都是零值，返回 -1 和 zero time。
func FindMinTime(ts []time.Time) (int, time.Time) {
	var min time.Time
	minIdx := -1
	for i, t := range ts {
		if t.IsZero() {
			continue // 跳过零值
		}
		if minIdx == -1 || t.Before(min) {
			min = t
			minIdx = i
		}
	}

	return minIdx, min
}
