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

// Incentive 时间激励函数：时间必须为第一参数（时序函数规范）。
// 输入当前时刻与系统状态（总成绩、在校人数、参考人数、达线人数），输出领导层感知的政绩标量。
// 实现：先按公式1算平均分（totalScore/studentCount），按公式2算升学率（examCount>0 时为 enrollCount/examCount，否则为 0），再调用 TeacherPayoff 得到政绩。
func Incentive(t time.Time, totalScore float64, studentCount int, examCount int, enrollCount int) float64 {
	if studentCount == 0 {
		return 0
	}
	avgScore := totalScore / float64(studentCount)
	var enrollRate float64
	if examCount > 0 {
		enrollRate = float64(enrollCount) / float64(examCount)
	}
	return TeacherPayoff(avgScore, enrollRate)
}

// IncentiveAt 与 Incentive 等价，仅强调时间第一参数的命名（时间序列函数规范）。
func IncentiveAt(t time.Time, totalScore float64, studentCount int, examCount int, enrollCount int) float64 {
	return Incentive(t, totalScore, studentCount, examCount, enrollCount)
}

// ========== 收益函数（Payoff） ==========

// TeacherPayoff 教师收益函数：以学生平均成绩和本科升学率为自变量。
// 公式：U_teacher = w_avg×平均成绩 + w_enroll×本科升学率，实现中 w_avg=0.6、w_enroll=0.4；政绩与教师收益同构。
func TeacherPayoff(avgScore, enrollRate float64) float64 {
	return 0.6*avgScore + 0.4*enrollRate
}

// GaokaoScore 学生高考成绩预测：与过往 3 年年末成绩正相关，近期权重更大。
// 公式：G = w1×S_{t-3} + w2×S_{t-2} + w3×S_{t-1}，w1+w2+w3=1；实现中 w1=0.2、w2=0.3、w3=0.5。
func GaokaoScore(scoreHistory [3]float64) float64 {
	const w1, w2, w3 = 0.2, 0.3, 0.5
	return w1*scoreHistory[0] + w2*scoreHistory[1] + w3*scoreHistory[2]
}

// StudentPayoff 学生收益函数：以个人高考成绩（预测值 GaokaoScore）作为收益；未在高考参考池则收益为 0。
func StudentPayoff(gaokaoScore float64, inExamPool bool) float64 {
	if !inExamPool {
		return 0
	}
	return gaokaoScore
}
