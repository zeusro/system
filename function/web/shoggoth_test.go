package web

import "testing"

func TestDynamicFindGirlfriends(t *testing.T) {
	// 初始化
	subjects := []*YoungAndBeautiful{
		{"妹子A", 0.9, 0.5, 1},
		{"妹子B", 0.7, 0.001, 2},
		{"妹子C", 0.8, 0.25, 3},
		{"妹子D", 0.85, 0.1, 4},
		{"妹子E", 0.75, 0.2, 5},
	}
	shoggoth := Shoggoth{TotalDays: 30, DailyTime: 24.0}
	shoggoth.DynamicFindGirlfriends(subjects)
}
