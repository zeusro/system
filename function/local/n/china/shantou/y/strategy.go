package y

import (
	"math/rand"
	"time"
)

// ChooseStrategy 根据角色与当前状态选择策略；时间 t 为第一参数（时序函数规范）。
// 按 a.Role 分派到对应角色的私有策略函数，返回该步选中的策略；黑曼巴在 Run 中不参与选策略，此处若被调用返回 StrategyNone。
func ChooseStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	switch a.Role {
	case RoleTeacherY:
		return chooseTeacherYStrategy(t, a, ctx)
	case RoleTeacherF:
		return chooseTeacherFStrategy(t, a, ctx)
	case RoleTeacherD:
		return chooseTeacherDStrategy(t, a, ctx)
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

// chooseTeacherYStrategy 教师 Y 的策略：以平均分为主；风险高则撒谎躲避，上期背叛则合作，平均分低且人多则 PUA。
// 实现顺序：LegalMoralRisk>0.6 → 撒谎躲避；上期背叛且剩余步数>1 → 正常教学；AvgScore<0.5 且 StudentCount>5 → PUA；否则正常教学。
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

// chooseTeacherFStrategy 教师 F 的策略：平均分+升学率导向；风险高则撒谎，上期背叛则合作，升学率低则减少参考人数。
// 实现顺序：LegalMoralRisk>0.5 → 撒谎躲避；上期背叛且剩余步数>1 → 正常教学；ExamCount>3 且 EnrollRate<0.6 → 减少参考人数；否则正常教学。
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

// chooseTeacherDStrategy 教师 D 的策略：以公式1+2为依据、平均成绩最大化，但不采取剔除学生策略（无 PUA、无减少参考人数）。
// 实现：LegalMoralRisk>0.5 → 撒谎躲避；否则始终正常教学。
func chooseTeacherDStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	if a.Factor.LegalMoralRisk > 0.5 {
		return StrategyLieEvade
	}
	return StrategyNormalTeach
}

// chooseStudentBlackMambaStrategy 学生黑曼巴的策略：钞能力、不参加高考，无策略、不参与博弈（y.md）。
// 实现：始终返回 StrategyNone；Run 中会跳过黑曼巴不参与选策略。
func chooseStudentBlackMambaStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	return StrategyNone
}

// chooseStudentJudasStrategy 学生犹大的策略：攀附教师 F；终局或上期教师背叛时采取网络暴力，否则努力学习。
// 实现：若 StudentCount>4 且 LegalMoralRisk<0.5，且（剩余步数≤1 或上期教师背叛）→ 网络暴力；否则努力学习。
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
// 实现：净 PUA 压力 = PUAExposure×(1-PUAResistance)；若净 PUA>0.5 且 Stress>0.6 → 休学；否则回避对抗。
func chooseStudentPStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	// 低 IQ、高 PUA 暴露且抵抗力低 → 休学消极对抗
	netPUA := a.Factor.PUAExposure * (1 - a.Factor.PUAResistance)
	if netPUA > 0.5 && a.Stress > 0.6 {
		return StrategyDropout
	}
	return StrategyAvoid
}

// chooseStudentYStrategy 学生 Y 的策略：高 IQ 运动员，依赖运动员加分；获取加分需学校领导同意（Run 中领导选「设计激励」即批准）。
// 实现：始终返回 StrategyAthleteBonus；加分是否生效由 Run 中 leaderApprovedBonus 决定。
func chooseStudentYStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	return StrategyAthleteBonus
}

// chooseStudentC13Strategy 学生 C13 的策略：高 IQ 贫困生，依赖努力升学，稳定选择努力学习。
// 实现：始终返回 StrategyStudyHard。
func chooseStudentC13Strategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	return StrategyStudyHard
}

// chooseGenericStudentStrategy 普通学生：按因子与压力在努力学习/回避对抗/休学间选择；上期教师背叛时倾向回避对抗。
// 实现：净 PUA>0.5 且 Stress>0.6 → 休学；上期背叛且剩余步数>2 → 回避对抗；IQ≥0.6 → 努力学习；否则回避对抗。
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
// 实现：AvgStress>0.4 → 减压安抚；否则 StrategyNone。
func choosePsychologistStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	if ctx.AvgStress > 0.4 {
		return StrategyDecompress
	}
	return StrategyNone
}

// chooseSchoolLeaderStrategy 学校领导的策略：政绩低于阈值时向下施压（踢猫），否则设计激励函数；设计激励即批准加分、安排心理老师定向辅导。
// 实现：调用 Incentive(t,...) 得到政绩；若政绩<0.5 → 向下施压；否则设计激励。
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

// ApplyStrategy 在时刻 t 根据策略 s 计算后果，返回 Consequence；不直接修改 Agent，由 Run 中根据后果写回状态。
// 实现：按策略类型填充 DeltaScore/DeltaStress/Dropout/LeaveExam/LegalRisk；PUA/减压等数值为模型参数（如 PUA 压力 0.3~0.7 随机，努力学习加分与 IQ 相关）。
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
