package model

import (
	"fmt"
	"math"
	"math/rand/v2"
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
