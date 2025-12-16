package problems

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

var (
	rng  *rand.Rand
	once sync.Once
)

func initRand() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt64_1_10M
// prompt：创建一个函数，随机生成1~10000000 的随机int64,要求每次调用生成的值都尽可能随机
func RandomInt64_1_10M() int64 {
	//•	只初始化一次随机种子
	//•	使用 crypto/rand 或高质量的 math/rand 初始化
	//•	每次调用只取随机值，不重复 seed
	once.Do(initRand)
	return rng.Int63n(10_000_000) + 1
}

func TestCollatz(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := RandomInt64_1_10M()
		// 测试角谷收敛性
		toOneSteps := ToOne(n)
		if toOneSteps != 1 {
			t.Errorf("ToOne(%d) returned non-positive steps: %d", n, toOneSteps)
		}
	}
}
