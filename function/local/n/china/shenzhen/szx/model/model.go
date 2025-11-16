package model

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

// 定义点结构
type Point struct {
	X float64
	Y float64
}

// Randomize 随机生成一个点，X、Y ∈ (0, 1000)，保留两位小数，分布尽可能离散
func RandonPoint() Point {
	p := Point{}
	p.X = math.Round(rand.Float64()*1000*100) / 100
	p.Y = math.Round(rand.Float64()*1000*100) / 100
	if p.X == 0 {
		p.X = 0.01
	}
	if p.Y == 0 {
		p.Y = 0.01
	}
	return p
}

func (p Point) String() string {
	return fmt.Sprintf("(%f,%f)", p.X, p.Y)
}

func (p Point) Compare(p2 Point) bool {
	return p.X == p2.X && p.Y == p2.Y
}

type Line struct {
	// Time time.Time
	A Point
	B Point
}

type Bean struct {
	Line
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

	// 将距离映射到 1ms~1000ms 之间（对数映射让增长更平滑）
	ms := 1 + int64(999*math.Tanh(dist/10)) // 距离大时趋近于1000ms

	// 加上 ±10% 的随机扰动
	jitter := rand.Float64()*0.2 - 0.1
	ms = int64(float64(ms) * (1 + jitter))

	if ms < 1 {
		ms = 1
	} else if ms > 1000 {
		ms = 1000
	}

	return time.Duration(ms) * time.Millisecond
}

type Aliyun struct {
}
type Alipay struct {
}

type PacMan interface {
	EatBeans(beans []Bean) map[time.Time]Bean
	GetCost() time.Duration
}
