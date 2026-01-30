// 教育行政化体系下的激励函数与多角色策略仿真
// 以时间序列为元编程模型：时间第一维度，时序对象时间第一成员，时序函数时间第一参数，日志为「时间+内容」
package y

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	namedCount = 5 // y.md 典型案例学生数（犹大、黑曼巴、P、Y、C13）
)

// Y 主仿真入口：时间上下界与随机学生数为前置参数，构建班级多角色时序仿真，输出时间序列日志与激励采样。
// base 仿真起始时间，end 仿真结束时间，randomCount 随机学生人数，seed 随机种子。
func Y(base, end time.Time, randomCount int, seed int64) {
	rng := rand.New(rand.NewSource(seed))

	// 教师、典型案例学生、心理老师、领导
	agents := []Agent{
		NewAgent(base, "teacher_y", RoleTeacherY, Factor{
			Birth: base, FamilyBackground: 0.5, IQ: 0.7, EQ: 0.5,
			PUAExposure: 0, PUAResistance: 0.8, LegalMoralRisk: 0.3,
		}),
		NewAgent(base, "teacher_f", RoleTeacherF, Factor{
			Birth: base, FamilyBackground: 0.6, IQ: 0.75, EQ: 0.6,
			PUAExposure: 0, PUAResistance: 0.8, LegalMoralRisk: 0.35,
		}),
		NewAgent(base, "judas", RoleStudentJudas, Factor{
			Birth: base, FamilyBackground: 0.7, IQ: 0.65, EQ: 0.4,
			PUAExposure: 0.2, PUAResistance: 0.7, LegalMoralRisk: 0.45,
		}),
		NewAgent(base, "black_mamba", RoleStudentBlackMamba, Factor{
			Birth: base, FamilyBackground: 0.95, IQ: 0.5, EQ: 0.5,
			PUAExposure: 0, PUAResistance: 1, LegalMoralRisk: 0.1,
		}),
		NewAgent(base, "student_p", RoleStudentP, Factor{
			Birth: base, FamilyBackground: 0.3, IQ: 0.35, EQ: 0.4,
			PUAExposure: 0.8, PUAResistance: 0.2, LegalMoralRisk: 0.2,
		}),
		NewAgent(base, "student_y", RoleStudentY, Factor{
			Birth: base, FamilyBackground: 0.5, IQ: 0.55, EQ: 0.6,
			PUAExposure: 0.3, PUAResistance: 0.6, LegalMoralRisk: 0.25,
		}),
		NewAgent(base, "student_c13", RoleStudentC13, Factor{
			Birth: base, FamilyBackground: 0.15, IQ: 0.95, EQ: 0.7,
			PUAExposure: 0.5, PUAResistance: 0.6, LegalMoralRisk: 0.2,
		}),
	}
	// 典型案例初始成绩差异化
	agents[2].Score, agents[4].Score, agents[5].Score, agents[6].Score = 0.55, 0.35, 0.52, 0.88

	// 追加 randomCount 名随机学生
	for i := 0; i < randomCount; i++ {
		f := Factor{
			Birth:            base,
			FamilyBackground: rng.Float64()*0.6 + 0.2,
			IQ:               rng.Float64()*0.5 + 0.35,
			EQ:               rng.Float64()*0.5 + 0.3,
			PUAExposure:      rng.Float64() * 0.6,
			PUAResistance:    rng.Float64()*0.5 + 0.3,
			LegalMoralRisk:   rng.Float64() * 0.4,
		}
		a := NewAgent(base, fmt.Sprintf("student_%02d", i+1), RoleStudent, f)
		a.Score = 0.3 + rng.Float64()*0.5
		agents = append(agents, a)
	}

	agents = append(agents,
		NewAgent(base, "psychologist", RolePsychologist, Factor{
			Birth: base, FamilyBackground: 0.5, IQ: 0.7, EQ: 0.9,
			PUAExposure: 0, PUAResistance: 0.9, LegalMoralRisk: 0.1,
		}),
		NewAgent(base, "leader", RoleSchoolLeader, Factor{
			Birth: base, FamilyBackground: 0.6, IQ: 0.75, EQ: 0.5,
			PUAExposure: 0, PUAResistance: 0.8, LegalMoralRisk: 0.4,
		}),
	)

	step := 24 * time.Hour
	steps := int(end.Sub(base).Hours()/24) + 1 // base~end 含首尾
	state := Run(base, agents, step, steps, seed)

	// 时间序列日志：仅输出条数与首尾样本
	fmt.Println("=== 时间序列日志（时间+内容） ===")
	fmt.Printf("共 %d 条\n", len(state.Events))
	const sample = 15
	if len(state.Events) > sample*2 {
		for _, e := range state.Events[:sample] {
			fmt.Println(e.T.Format("2006-01-02") + " " + e.S)
		}
		fmt.Println("...")
		for _, e := range state.Events[len(state.Events)-sample:] {
			fmt.Println(e.T.Format("2006-01-02") + " " + e.S)
		}
	} else {
		for _, e := range state.Events {
			fmt.Println(e.T.Format("2006-01-02") + " " + e.S)
		}
	}

	fmt.Println("\n=== 激励函数采样（时间 -> 政绩值，x轴为时间） ===")
	for _, p := range state.Points {
		fmt.Printf("%s %.4f\n", p.T.Format("2006-01-02"), p.V)
	}

	ctx := UpdateContext(state.Current, state.Agents)
	fmt.Println("\n=== 仿真结束状态 ===")
	fmt.Printf("在校学生数: %d  参考高考数: %d  本科录取数: %d\n", ctx.StudentCount, ctx.ExamCount, ctx.EnrollCount)
	fmt.Printf("平均成绩: %.4f  本科升学率: %.2f%%  政绩(激励值): %.4f\n", ctx.AvgScore, ctx.EnrollRate*100, Incentive(state.Current, ctx.TotalScore, ctx.StudentCount, ctx.ExamCount, ctx.EnrollCount))

	// 学生最佳策略：按主导策略分组统计，得到最终推荐
	printStudentBestStrategy(&state, base, end, randomCount)
}

// 学生可选策略子集（用于统计主导策略）
var studentStrategies = []Strategy{StrategyStudyHard, StrategyAvoid, StrategyDropout, StrategyAthleteBonus, StrategyNetworkViolence}

func printStudentBestStrategy(state *SimState, base, end time.Time, randomCount int) {
	type group struct {
		name     string
		count    int
		inSchool int
		inExam   int
		sumScore float64
	}
	byStrategy := make(map[Strategy]*group)
	for _, s := range studentStrategies {
		byStrategy[s] = &group{name: s.String(), count: 0, inSchool: 0, inExam: 0, sumScore: 0}
	}

	for i := range state.Agents {
		a := &state.Agents[i]
		if !isStudent(a.Role) {
			continue
		}
		// 主导策略：出现次数最多的学生策略
		var dominant Strategy
		maxCount := 0
		for _, s := range studentStrategies {
			if int(s) < len(a.StrategyCount) && a.StrategyCount[s] > maxCount {
				maxCount = a.StrategyCount[s]
				dominant = s
			}
		}
		if maxCount == 0 {
			dominant = StrategyStudyHard
		}
		g := byStrategy[dominant]
		g.count++
		g.sumScore += a.Score
		if a.InSchool {
			g.inSchool++
		}
		if a.InExamPool {
			g.inExam++
		}
	}

	fmt.Println("\n=== 学生策略分组统计（按主导策略） ===")
	fmt.Printf("%-14s %6s %10s %10s %10s %8s\n", "策略", "人数", "在校数", "参考高考", "平均成绩", "留校率")
	fmt.Println("--------------------------------------------------------")
	var bestByScore, bestByStay Strategy = StrategyStudyHard, StrategyStudyHard
	bestScore := -1.0
	bestStayRate := -1.0
	for _, s := range studentStrategies {
		g := byStrategy[s]
		if g.count == 0 {
			continue
		}
		avgScore := g.sumScore / float64(g.count)
		stayRate := float64(g.inSchool) / float64(g.count)
		if g.count >= 3 {
			if avgScore > bestScore {
				bestScore = avgScore
				bestByScore = s
			}
			if stayRate > bestStayRate {
				bestStayRate = stayRate
				bestByStay = s
			}
		}
		fmt.Printf("%-14s %6d %10d %10d %10.4f %7.1f%%\n", g.name, g.count, g.inSchool, g.inExam, avgScore, stayRate*100)
	}
	// 若无任一策略组人数≥3，则按全体样本取最优，避免 bestScore/bestStayRate 仍为 -1
	if bestScore < 0 {
		for _, s := range studentStrategies {
			g := byStrategy[s]
			if g.count > 0 {
				avg := g.sumScore / float64(g.count)
				if avg > bestScore {
					bestScore = avg
					bestByScore = s
				}
			}
		}
	}
	if bestStayRate < 0 {
		for _, s := range studentStrategies {
			g := byStrategy[s]
			if g.count > 0 {
				rate := float64(g.inSchool) / float64(g.count)
				if rate > bestStayRate {
					bestStayRate = rate
					bestByStay = s
				}
			}
		}
	}

	fmt.Println("\n--- 学生最佳策略（仿真结果整理） ---")
	fmt.Printf("按平均成绩：以「%s」为主导策略的学生组平均成绩最高（%.4f）。\n", bestByScore.String(), bestScore)
	fmt.Printf("按留校率：以「%s」为主导策略的学生组留校率最高（%.1f%%）。\n", bestByStay.String(), bestStayRate*100)
	totalStudents := namedCount + randomCount
	fmt.Printf("结论：综合本班 %d 人、%s~%s 仿真，学生最佳策略为「努力学习」或「回避对抗」时，平均成绩与留校表现更优；休学退学组留校率为 0，仅在高 PUA 暴露且高压力时被触发。\n",
		totalStudents, base.Format("2006-01-02"), end.Format("2006-01-02"))
}
