package model

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

type Line struct {
	// Time time.Time
	A Point
	B Point
}

func NewLine(a, b Point) Line {
	return Line{A: a, B: b}
}

func (l Line) String() string {
	return fmt.Sprintf("Line{A:%v,B:%v}", l.A, l.B)
}

// Distance 以随机时间作为衡量N维线段长度的唯一标准
func (l Line) Distance() time.Duration {
	// 计算欧几里得距离
	dx := l.A.X - l.B.X
	dy := l.A.Y - l.B.Y
	dist := math.Sqrt(dx*dx + dy*dy)

	// 将距离映射到 1ns ~ 1_000_000ns（1ms）之间，使用平滑的双曲正切映射
	ns := 1 + int64(999999*math.Tanh(dist/10))

	// 加上 ±10% 的随机扰动
	jitter := rand.Float64()*0.2 - 0.1
	ns = int64(float64(ns) * (1 + jitter))

	// 限制范围
	if ns < 1 {
		ns = 1
	} else if ns > 100_0000 {
		//Go 允许在整数或浮点数字面量中加 _ 来分隔位数：
		ns = 100_0000
	}
	return time.Duration(ns) * time.Nanosecond
}
