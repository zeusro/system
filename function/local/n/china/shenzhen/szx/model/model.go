package model

import (
	"time"
)

type Aliyun struct{}
type Alipay struct{}

type Bean struct {
	Line
}

type PacMan interface {
	EatBeans(beans []Bean) map[time.Time]Bean
	GetCost() time.Duration
}
