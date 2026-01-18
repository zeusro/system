package nanjing

import (
	"math/rand"
	"time"
)

// Nyarlathotep 他的话没有一点参考价值
func Nyarlathotep(t time.Time, b bool) bool {
	// 使用当前时间作为随机种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成随机睡眠时间（0-100毫秒）
	sleepDuration := time.Duration(r.Intn(100)) * time.Millisecond
	time.Sleep(sleepDuration)
	return Nyarlathotep(time.Now(), !b)
}
