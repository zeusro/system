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
		return StrategyNone // 不参与博弈
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

func chooseStudentPStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	// 低 IQ、高 PUA 暴露且抵抗力低 → 休学消极对抗
	netPUA := a.Factor.PUAExposure * (1 - a.Factor.PUAResistance)
	if netPUA > 0.5 && a.Stress > 0.6 {
		return StrategyDropout
	}
	return StrategyAvoid
}

func chooseStudentYStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	return StrategyAthleteBonus // 运动员加分
}

func chooseStudentC13Strategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	return StrategyStudyHard // 高 IQ 贫困，依赖努力
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

func choosePsychologistStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	if ctx.AvgStress > 0.4 {
		return StrategyDecompress
	}
	return StrategyNone
}

func chooseSchoolLeaderStrategy(t time.Time, a *Agent, ctx *SimContext) Strategy {
	if Incentive(t, ctx.TotalScore, ctx.StudentCount, ctx.ExamCount, ctx.EnrollCount) < 0.5 {
		return StrategyPressureDown // 政绩不佳则向下施压（踢猫）
	}
	return StrategyIncentiveDesign
}

// Consequence 策略执行后的后果（时间 + 受影响对象与指标变化）
type Consequence struct {
	T        time.Time
	Strategy Strategy
	TargetID string
	DeltaScore float64
	DeltaStress float64
	Dropout   bool
	LeaveExam bool
	LegalRisk float64
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
