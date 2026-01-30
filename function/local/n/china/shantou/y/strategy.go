package y

import (
	"math/rand"
	"time"
)

// ChooseStrategy 根据角色与当前状态选择策略（时间第一参数）
func ChooseStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	switch a.Role {
	case RoleTeacherY:
		return chooseTeacherYStrategy(t, a, ctx)
	case RoleTeacherF:
		return chooseTeacherFStrategy(t, a, ctx)
	case RoleStudentJudas:
		return chooseStudentJudasStrategy(t, a, ctx)
	case RoleStudentBlackMamba:
		return chooseStudentBlackMambaStrategy(t, a, ctx)
	case RoleStudentP:
		return chooseStudentPStrategy(t, a, ctx)
	case RoleStudentY:
		return chooseStudentYStrategy(t, a, ctx)
	case RoleStudentC13:
		return chooseStudentC13Strategy(t, a, ctx)
	case RoleStudent:
		return chooseGenericStudentStrategy(t, a, ctx)
	case RolePsychologist:
		return choosePsychologistStrategy(t, a, ctx)
	case RoleSchoolLeader:
		return chooseSchoolLeaderStrategy(t, a, ctx)
	default:
		return StrategyNone
	}
}

// chooseTeacherYStrategy 教师 Y 的策略：以平均分为主，风险高则撒谎躲避，上期背叛则合作，平均分低且人多则 PUA。
func chooseTeacherYStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	if a.Factor.LegalMoralRisk > 0.6 {
		return StrategyLieEvade
	}
	// 重复博弈：上期已背叛则本期合作以恢复声誉
	if ctx.LastRoundTeacherDefection && ctx.StepsRemaining > 1 {
		return StrategyNormalTeach
	}
	if ctx.AvgScore < 0.5 && ctx.StudentCount > 5 {
		return StrategyPUA
	}
	return StrategyNormalTeach
}

// chooseTeacherFStrategy 教师 F 的策略：平均分+升学率导向，风险高则撒谎，上期背叛则合作，升学率低则减少参考人数。
func chooseTeacherFStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	if a.Factor.LegalMoralRisk > 0.5 {
		return StrategyLieEvade
	}
	if ctx.LastRoundTeacherDefection && ctx.StepsRemaining > 1 {
		return StrategyNormalTeach
	}
	if ctx.ExamCount > 3 && ctx.EnrollRate < 0.6 {
		return StrategyReduceExamCount
	}
	return StrategyNormalTeach
}

// chooseStudentBlackMambaStrategy 学生黑曼巴的策略：钞能力，可不参加高考（y.md：家里非常有钱，可直接购买澳门科技大学本科学位），无策略、不参与博弈。
func chooseStudentBlackMambaStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	return StrategyNone
}

// chooseStudentJudasStrategy 学生犹大的策略：攀附教师 F，终局或上期教师背叛时采取网络暴力，否则努力学习。
func chooseStudentJudasStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	// 重复博弈：终局时背叛；中期若教师上期背叛则可报复（网络暴力）
	if ctx.StudentCount > 4 && a.Factor.LegalMoralRisk < 0.5 {
		if ctx.StepsRemaining <= 1 {
			return StrategyNetworkViolence
		}
		if ctx.LastRoundTeacherDefection {
			return StrategyNetworkViolence
		}
	}
	return StrategyStudyHard
}

// chooseStudentPStrategy 学生 P 的策略：低 IQ、高 PUA 净暴露且压力大时休学，否则回避对抗。
func chooseStudentPStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	// 低 IQ、高 PUA 暴露且抵抗力低 → 休学消极对抗
	netPUA := a.Factor.PUAExposure * (1 - a.Factor.PUAResistance)
	if netPUA > 0.5 && a.Stress > 0.6 {
		return StrategyDropout
	}
	return StrategyAvoid
}

// chooseStudentYStrategy 学生 Y 的策略：高IQ运动员，依赖运动员加分；获取加分需学校领导同意（Run 中领导选择「设计激励」即批准）。
func chooseStudentYStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	return StrategyAthleteBonus
}

// chooseStudentC13Strategy 学生 C13 的策略：高 IQ 贫困生，依赖努力升学，始终选择努力学习。
func chooseStudentC13Strategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	return StrategyStudyHard
}

// chooseGenericStudentStrategy 普通学生：按因子与压力在 努力学习/回避对抗/休学 间选择
// 重复博弈：上期教师背叛时倾向回避对抗（减少暴露于 PUA）
func chooseGenericStudentStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	netPUA := a.Factor.PUAExposure * (1 - a.Factor.PUAResistance)
	if netPUA > 0.5 && a.Stress > 0.6 {
		return StrategyDropout
	}
	if ctx.LastRoundTeacherDefection && ctx.StepsRemaining > 2 {
		return StrategyAvoid
	}
	if a.Factor.IQ >= 0.6 {
		return StrategyStudyHard
	}
	return StrategyAvoid
}

// choosePsychologistStrategy 心理老师的策略：全体平均压力高于阈值时减压安抚，否则不行动。
func choosePsychologistStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	if ctx.AvgStress > 0.4 {
		return StrategyDecompress
	}
	return StrategyNone
}

// chooseSchoolLeaderStrategy 学校领导的策略：负责分配资源、安排心理老师定向辅导；政绩低于阈值时向下施压（踢猫），否则设计激励函数。
func chooseSchoolLeaderStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	if Incentive(t, ctx.TotalScore, ctx.StudentCount, ctx.ExamCount, ctx.EnrollCount) < 0.5 {
		return StrategyPressureDown
	}
	return StrategyIncentiveDesign
}

// Consequence 策略执行后的后果（时间 + 受影响对象与指标变化），供 Run 用于更新 Agent 与事件日志。
type Consequence struct {
	T         time.Time  // 发生时刻
	Strategy  Strategy   // 所采取的策略
	TargetID  string     // 受影响目标 ID（若适用）
	DeltaScore  float64  // 成绩变化量
	DeltaStress float64  // 压力变化量
	Dropout    bool     // 是否触发休学/退学
	LeaveExam  bool     // 是否退出高考参考
	LegalRisk  float64  // 法规道德风险增量
}

// ApplyStrategy 在时刻 t 应用策略，更新 Agent 与上下文，返回后果
func ApplyStrategy(t time.Time, s Strategy, a *Agent, ctx *SimContext, rng *rand.Rand) Consequence {
	c := Consequence{T: t, Strategy: s}
	switch s {
	case StrategyPUA:
		// 随机增加某学生压力，可能触发休学
		c.DeltaStress = 0.3 + 0.4*rng.Float64()
		c.LegalRisk = 0.1
	case StrategyReduceExamCount:
		c.LeaveExam = true
		c.LegalRisk = 0.15
	case StrategyLieEvade:
		c.LegalRisk = -0.05 // 短期“降低”可见风险
	case StrategyDropout:
		c.Dropout = true
		c.LeaveExam = true
	case StrategyNetworkViolence:
		c.DeltaStress = 0.4
		c.LegalRisk = 0.2
	case StrategyDecompress:
		c.DeltaStress = -0.2
	case StrategyPressureDown:
		c.DeltaStress = 0.25 // 向下传递压力
	case StrategyAthleteBonus:
		c.DeltaScore = 0.08
	case StrategyStudyHard:
		c.DeltaScore = 0.05 + 0.05*a.Factor.IQ
	case StrategyAvoid:
		c.DeltaStress = -0.05
	default:
		return c
	}
	return c
}
