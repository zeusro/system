package y

import (
	"fmt"
	"math/rand"
	"time"
)

// SimContext 当前时刻的系统聚合状态（用于策略与激励计算）
// 重复博弈扩展：StepsRemaining 与 LastRoundTeacherDefection 用于历史/视界依赖策略。
type SimContext struct {
	Now          time.Time
	TotalScore   float64
	StudentCount int
	ExamCount    int
	EnrollCount  int
	AvgScore     float64
	EnrollRate   float64
	AvgStress    float64

	StepsRemaining           int  // 剩余步数（用于终局效应与重复博弈视界）
	LastRoundTeacherDefection bool // 上期是否有教师采取 PUA/减参考人数（用于报复/合作）
}

// SimState 时间序列空间中的仿真状态（时间第一成员）
type SimState struct {
	Birth    time.Time // 仿真起始时刻
	Current  time.Time
	Agents   []Agent
	Events   []Event  // 时间+内容 日志
	Points   []Point  // 激励函数随时间采样
	Duration time.Duration
}

func isStudent(r Role) bool {
	return r == RoleStudentJudas || r == RoleStudentBlackMamba || r == RoleStudentP || r == RoleStudentY || r == RoleStudentC13 || r == RoleStudent
}

func clampAgentState(a *Agent) {
	if a.Score > 1 {
		a.Score = 1
	}
	if a.Score < 0 {
		a.Score = 0
	}
	if a.Stress > 1 {
		a.Stress = 1
	}
	if a.Stress < 0 {
		a.Stress = 0
	}
}

// LogTS 时间序列日志：打印/记录必须为「时间+内容」格式
func LogTS(t time.Time, format string, args ...interface{}) Event {
	content := fmt.Sprintf(format, args...)
	return Event{T: t, S: t.Format(time.RFC3339) + " " + content}
}

// UpdateContext 根据当前 Agent 列表更新 SimContext
func UpdateContext(now time.Time, agents []Agent) SimContext {
	var totalScore float64
	var totalStress float64
	nStudent, nExam, nEnroll := 0, 0, 0
	for i := range agents {
		a := &agents[i]
		if a.Role == RoleSchoolLeader || a.Role == RolePsychologist {
			continue
		}
		if a.Role == RoleTeacherY || a.Role == RoleTeacherF {
			continue
		}
		if !a.InSchool {
			continue
		}
		nStudent++
		totalScore += a.Score
		totalStress += a.Stress
		if a.InExamPool {
			nExam++
			if a.Score >= 0.6 { // 简化为分数达线即录取
				nEnroll++
			}
		}
	}
	avgScore, enrollRate := 0.0, 0.0
	if nStudent > 0 {
		avgScore = totalScore / float64(nStudent)
		avgStress := totalStress / float64(nStudent)
		totalStress = avgStress
	}
	if nExam > 0 {
		enrollRate = float64(nEnroll) / float64(nExam)
	}
	return SimContext{
		Now:          now,
		TotalScore:   totalScore,
		StudentCount: nStudent,
		ExamCount:    nExam,
		EnrollCount:  nEnroll,
		AvgScore:     avgScore,
		EnrollRate:   enrollRate,
		AvgStress:    totalStress,
	}
}

// Run 运行时间步进仿真；步长 step，步数 steps。
// 重复博弈：每步先全体选策略（基于剩余步数及上期背叛），再统一施加后果，避免顺序依赖。
func Run(birth time.Time, agents []Agent, step time.Duration, steps int, seed int64) SimState {
	rng := rand.New(rand.NewSource(seed))
	state := SimState{Birth: birth, Current: birth, Agents: agents, Duration: step * time.Duration(steps)}
	chosen := make([]Strategy, len(agents))
	lastRoundTeacherDefection := false

	for i := 0; i < steps; i++ {
		now := birth.Add(step * time.Duration(i))
		state.Current = now
		ctx := UpdateContext(now, state.Agents)
		ctx.StepsRemaining = steps - i
		ctx.LastRoundTeacherDefection = lastRoundTeacherDefection

		// 采样激励函数（时间序列可视化用）
		inc := Incentive(now, ctx.TotalScore, ctx.StudentCount, ctx.ExamCount, ctx.EnrollCount)
		state.Points = append(state.Points, Point{T: now, V: inc})

		// 阶段一：全体在册成员同时选策略（重复博弈：基于同一 ctx 与历史）
		for j := range state.Agents {
			chosen[j] = StrategyNone
			a := &state.Agents[j]
			if !a.InSchool {
				continue
			}
			// 学生黑曼巴：钞能力、不参与高考，不参与策略博弈（y.md）
			if a.Role == RoleStudentBlackMamba {
				continue
			}
			chosen[j] = ChooseStrategy(now, a, &ctx)
		}

		// 本步领导是否同意加分（y.md：学生 Y 获取加分需学校领导同意；领导选择「设计激励」即视为同意）
		leaderApprovedBonus := false
		for j := range state.Agents {
			if state.Agents[j].Role == RoleSchoolLeader && chosen[j] == StrategyIncentiveDesign {
				leaderApprovedBonus = true
				break
			}
		}

		// 阶段二：统一施加后果并更新 LastStrategy
		lastRoundTeacherDefection = false
		for j := range state.Agents {
			a := &state.Agents[j]
			s := chosen[j]
			if s == StrategyNone {
				continue
			}
			if len(a.StrategyCount) > int(s) {
				a.StrategyCount[s]++
			}
			a.LastStrategy = s
			c := ApplyStrategy(now, s, a, &ctx, rng)
			state.Events = append(state.Events, LogTS(now, "[%s] %s 采取策略 %s", a.ID, a.Role.String(), s.String()))
			a.LegalRisk += c.LegalRisk

			if s == StrategyPUA || s == StrategyReduceExamCount {
				lastRoundTeacherDefection = true
			}

			// 休学/退学：后果作用于行为者本人
			if c.Dropout && isStudent(a.Role) {
				a.InSchool = false
				a.InExamPool = false
				state.Events = append(state.Events, LogTS(now, "[后果] %s 休学/退学", a.ID))
			}
			// 运动员加分/努力学习/回避对抗：作用于本人（y.md：学生 Y 加分需领导同意）
			if (s == StrategyAthleteBonus || s == StrategyStudyHard || s == StrategyAvoid) && isStudent(a.Role) {
				deltaScore := c.DeltaScore
				if s == StrategyAthleteBonus && a.Role == RoleStudentY && !leaderApprovedBonus {
					deltaScore = 0
					state.Events = append(state.Events, LogTS(now, "[后果] %s 运动员加分需领导同意，本期未批准", a.ID))
				}
				a.Score += deltaScore
				a.Stress += c.DeltaStress
				clampAgentState(a)
			}
			// 减压：作用于全体学生（简化：选一目标）
			if s == StrategyDecompress {
				targetIdx := pickStudentTarget(&state.Agents, a.Role, rng)
				if targetIdx >= 0 {
					t := &state.Agents[targetIdx]
					t.Stress += c.DeltaStress
					clampAgentState(t)
				}
			}
			// PUA/网络暴力/向下施压/减少参考人数：作用于随机目标学生
			if (s == StrategyPUA || s == StrategyNetworkViolence || s == StrategyPressureDown || s == StrategyReduceExamCount) && (c.DeltaStress != 0 || c.LeaveExam) {
				targetIdx := pickStudentTarget(&state.Agents, a.Role, rng)
				if targetIdx >= 0 {
					t := &state.Agents[targetIdx]
					t.Stress += c.DeltaStress
					if c.LeaveExam {
						t.InExamPool = false
						state.Events = append(state.Events, LogTS(now, "[后果] %s 退出高考参考", t.ID))
					}
					clampAgentState(t)
				}
			}
		}
	}

	return state
}

func pickStudentTarget(agents *[]Agent, from Role, rng *rand.Rand) int {
	var candidates []int
	for i := range *agents {
		a := &(*agents)[i]
		if !a.InSchool {
			continue
		}
		switch a.Role {
		case RoleStudentJudas, RoleStudentBlackMamba, RoleStudentP, RoleStudentY, RoleStudentC13, RoleStudent:
			candidates = append(candidates, i)
		}
	}
	if len(candidates) == 0 {
		return -1
	}
	return candidates[rng.Intn(len(candidates))]
}
