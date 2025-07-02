package v3

import "testing"

func TestOnlyLoveYou(t *testing.T) {
	OnlyLoveYou()
}

func TestLoveYouAll111(t *testing.T) {
	LoveYouAll111()
}

// BenchmarkLoveYouAll111 用于性能测试 LoveYouAll111 的核心算法
// 要运行基准测试，请放在 _test.go 文件中并使用 go test -bench=. 命令运行
// 这里为了方便演示直接放在此文件，实际应放在 sea_eye_test.go
func BenchmarkLoveYouAll111(b *testing.B) {
	var limit int64 = 36000  // s
	var distance int64 = 600 // m

	// 创建多个测试对象用于模拟压力
	numObjects := 100
	stars := make([]NewSwallowGarden, 0, numObjects)
	for i := 1; i <= numObjects; i++ {
		stars = append(stars, NewSwallowGarden{
			Limit:    limit,
			Distance: distance,
			V:        int64(i + 1), // 保证速度不同
		})
	}
	sun := NewSwallowGarden{
		Limit:    limit,
		Distance: distance,
		V:        1, // 作为中心对象
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sun.NP(stars)
	}
}
