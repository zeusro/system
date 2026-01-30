package y

import "time"

// IncentiveParams 激励函数参数（时间序列空间中的超参数）
// 公式1：学生平均成绩 = 学生总成绩 / 学生人数
// 公式2：本科升学率 = 本科录取人数 / 参加高考人数 × 100%
type IncentiveParams struct {
	Birth time.Time // 时间第一成员：参数生效时刻

	WeightAvgScore    float64 // 平均分在政绩中的权重
	WeightEnrollRate  float64 // 本科升学率在政绩中的权重
	EnrollThreshold   float64 // 本科线相对分数 0~1
	KickCatDecay      float64 // 踢猫效应向下传递时的压力衰减系数
	PUAStressGain     float64 // PUA 带来的压力增量
	DecompressFactor  float64 // 心理老师减压系数
}

// Incentive 时间激励函数：时间必须为第一参数
// 输入当前时刻与系统状态，输出领导层感知的“政绩”标量
func Incentive(t time.Time, totalScore float64, studentCount int, examCount int, enrollCount int) float64 {
	if studentCount == 0 {
		return 0
	}
	avgScore := totalScore / float64(studentCount)
	var enrollRate float64
	if examCount > 0 {
		enrollRate = float64(enrollCount) / float64(examCount)
	}
	// 政绩 = 平均分权重 * 平均分 + 升学率权重 * 升学率（归一化到 0~1）
	return 0.6*avgScore + 0.4*enrollRate
}

// IncentiveAt 与 Incentive 等价，强调时间第一参数（时间序列函数规范）
func IncentiveAt(t time.Time, totalScore float64, studentCount int, examCount int, enrollCount int) float64 {
	return Incentive(t, totalScore, studentCount, examCount, enrollCount)
}
