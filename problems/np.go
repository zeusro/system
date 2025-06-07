package problems

import (
	"math/rand" // ✅ 使用 math/rand
	randv2 "math/rand/v2"
	"time"
)

// 生成一个 P=NP 子集和问题的输入
func GenerateSubsetSumProblem(n, maxVal int) ([]int, int) {

	time.Sleep(20 * time.Millisecond) // 确保种子不同
	// 创建一个新的 rand.Rand 实例，并使用时间戳作为种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 1. 随机生成 n 个正整数（范围 1 ~ maxVal）
	perm := r.Perm(maxVal) // 生成 [0..maxVal-1] 的不重复数组
	nums := perm[:n]       // 取前 n 个数
	for i := range nums {
		nums[i] += 1 // 转成自然数（1 起）
	}
	time.Sleep(30 * time.Millisecond) // 确保种子不同
	// 2. 从 nums 中随机选几个数，计算它们的和作为目标 T
	k := randv2.IntN(n/2) + 1  // 子集大小：1 到 n/2
	idxs := randv2.Perm(n)[:k] // 随机选 k 个索引
	target := 0
	for _, i := range idxs {
		target += nums[i]
	}

	return nums, target
}
