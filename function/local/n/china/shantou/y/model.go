// Package shantou 教育行政化体系下的激励函数与时间序列元编程模型
// 第一性原理：时间是第一维度。所有时序对象时间必须为第一成员。
package y

import "time"

// ========== 时间序列对象（时间第一成员） ==========

// Factor 量化因子：家庭背景、IQ、EQ、PUA暴露/抵抗力、法律法规道德风险
// 取值 [0,1] 归一化，或使用标准分
type Factor struct {
	Birth time.Time // 时间第一成员：因子生效/采样时刻

	// 家庭背景 FamilyBackground: 0=极贫 1=极富，影响资源与升学路径
	FamilyBackground float64
	// IQ: 0~1 映射到智力分数，影响学业表现与策略理解
	IQ float64
	// EQ: 0~1 情绪智力，影响抗压与踢猫链中的传播
	EQ float64
	// PUAExposure: 教师PUA策略对该个体的暴露强度 0~1
	PUAExposure float64
	// PUAResistance: 个体对PUA的抵抗力 0~1
	PUAResistance float64
	// LegalMoralRisk: 个体/行为触发的法律法规道德风险 0~1，越高越易被追责
	LegalMoralRisk float64
}

// Event 时序事件：时间 + 内容（满足时间序列日志规范）
type Event struct {
	T time.Time
	S string
}

// Point 时序空间中的点（时间 + 标量），用于激励函数与可视化
type Point struct {
	T time.Time
	V float64
}

// NLine 参数化线段（时间序列空间中的线段）
type NLine struct {
	Start time.Time
	End   time.Time
	V0    float64
	V1    float64
}
