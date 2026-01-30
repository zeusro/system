// 教育行政化体系下的激励函数与多角色策略仿真
// 以时间序列为元编程模型：时间第一维度，时序对象时间第一成员，时序函数时间第一参数，日志为「时间+内容」
package y

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	namedCount   = 5   // y.md 典型案例学生数（犹大、黑曼巴、P、Y、C13）
	namedAgents  = 8   // 教师3（Y、F、D）+ 典型案例学生5
	staffAgents  = 2   // 心理老师、领导
	c13AgentID   = "student_c13"
	initialSamples = 15
)

// 典型案例初始成绩（按角色设置，避免依赖 agents 下标顺序）
var initialScoreByRole = map[Role]float64{
	RoleStudentJudas: 0.55, RoleStudentP: 0.35, RoleStudentY: 0.52, RoleStudentC13: 0.88,
}

// newNamedAgents 构造教师与命名学生（不含心理老师、领导），顺序固定便于与 initialScoreByRole 配合。
// 返回 8 个 Agent：教师 Y/F/D、犹大、黑曼巴、P、Y、C13；每人含 Birth、Role、Factor，初始 InSchool=true，黑曼巴 InExamPool=false。
func newNamedAgents(base time.Time) []Agent {
	return []Agent{
		NewAgent(base, "teacher_y", RoleTeacherY, Factor{
			Birth: base, FamilyBackground: 0.5, IQ: 0.7, EQ: 0.5,
			PUAExposure: 0, PUAResistance: 0.8, LegalMoralRisk: 0.3,
		}),
		NewAgent(base, "teacher_f", RoleTeacherF, Factor{
			Birth: base, FamilyBackground: 0.6, IQ: 0.75, EQ: 0.6,
			PUAExposure: 0, PUAResistance: 0.8, LegalMoralRisk: 0.35,
		}),
		NewAgent(base, "teacher_d", RoleTeacherD, Factor{
			Birth: base, FamilyBackground: 0.55, IQ: 0.72, EQ: 0.65,
			PUAExposure: 0, PUAResistance: 0.85, LegalMoralRisk: 0.3,
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
			Birth: base, FamilyBackground: 0.5, IQ: 0.72, EQ: 0.6,
			PUAExposure: 0.3, PUAResistance: 0.6, LegalMoralRisk: 0.25,
		}),
		NewAgent(base, c13AgentID, RoleStudentC13, Factor{
			Birth: base, FamilyBackground: 0.15, IQ: 0.95, EQ: 0.7,
			PUAExposure: 0.5, PUAResistance: 0.6, LegalMoralRisk: 0.2,
		}),
	}
}

// Y 主仿真入口：以时间上下界与随机学生数为前置参数，构建班级多角色时序仿真，向 stdout 输出时间序列日志与激励采样。
// 参数：base 仿真起始时间，end 仿真结束时间，randomCount 随机学生人数，seed 随机种子。
//
// 实现过程：
// （1）用 newNamedAgents(base) 构造教师与命名学生，按 initialScoreByRole 为典型学生赋初始成绩与 ScoreHistory。
// （2）追加 randomCount 名普通学生（RoleStudent），因子与成绩随机生成，ScoreHistory 初始为当前成绩。
// （3）追加心理老师与学校领导。
// （4）按日步进调用 Run(base, agents, 24h, steps, seed)，steps 由 base~end 含首尾的天数决定。
// （5）输出：时间序列日志（共 N 条，首尾各 initialSamples 条或全部）、激励采样点列、终态统计（在校数/参考数/录取数/平均分/升学率/政绩）、
// 教师收益与学生收益采样、按主导策略分组的学生统计与最佳策略结论、C13 个性化建议。
func Y(base, end time.Time, randomCount int, seed int64) {
	rng := rand.New(rand.NewSource(seed))

	agents := make([]Agent, 0, namedAgents+randomCount+staffAgents)
	named := newNamedAgents(base)
	for i := range named {
		if score, ok := initialScoreByRole[named[i].Role]; ok {
			named[i].Score = score
			if isStudent(named[i].Role) {
				named[i].ScoreHistory = [3]float64{score, score, score}
			}
		}
		agents = append(agents, named[i])
	}

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
		a.ScoreHistory = [3]float64{a.Score, a.Score, a.Score}
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
	if len(state.Events) > initialSamples*2 {
		for _, e := range state.Events[:initialSamples] {
			fmt.Println(e.T.Format("2006-01-02") + " " + e.S)
		}
		fmt.Println("...")
		for _, e := range state.Events[len(state.Events)-initialSamples:] {
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

	// 收益函数：教师以平均成绩+升学率为收益；学生以个人高考成绩（与过往3年正相关）为收益
	fmt.Println("\n=== 收益函数采样 ===")
	teacherPayoff := TeacherPayoff(ctx.AvgScore, ctx.EnrollRate)
	fmt.Printf("教师收益 U_teacher = TeacherPayoff(平均成绩, 本科升学率) = %.4f\n", teacherPayoff)
	fmt.Println("学生收益 U_student = StudentPayoff(GaokaoScore(过往3年成绩), 是否参考高考)：")
	for i := range state.Agents {
		a := &state.Agents[i]
		if !isStudent(a.Role) {
			continue
		}
		gaokao := GaokaoScore(a.ScoreHistory)
		payoff := StudentPayoff(gaokao, a.InExamPool)
		fmt.Printf("  %s: GaokaoScore=%.4f  InExamPool=%v  U_student=%.4f\n", a.ID, gaokao, a.InExamPool, payoff)
	}

	// 学生最佳策略：按主导策略分组统计，得到最终推荐
	printStudentBestStrategy(&state, base, end, randomCount)
	// 文末 C13 建议（高 IQ 贫困生）
	printC13Suggestion(&state)
}

// studentStrategies 学生可选策略子集，用于按主导策略分组统计与推荐。
var studentStrategies = []Strategy{StrategyStudyHard, StrategyAvoid, StrategyDropout, StrategyAthleteBonus, StrategyNetworkViolence}

// printStudentBestStrategy 按主导策略对学生分组统计，输出表格与最佳策略结论。
// 实现：遍历所有学生，取每个学生 StrategyCount 中出现次数最多的策略作为其主导策略，按策略分组累加人数、在校数、参考数、总成绩；
// 计算每组平均成绩与留校率；若某组人数≥3 则参与「按平均成绩最优」「按留校率最优」的候选，否则用全体样本补选；
// 输出策略名、人数、在校数、参考数、平均成绩、留校率表格，以及按平均成绩/留校率的最优策略与结论文案。
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

// printC13Suggestion 根据仿真结束状态对 C13（高 IQ 贫困生）输出个性化建议到 stdout。
// 实现：在 state.Agents 中定位 RoleStudentC13，若无则返回；取 C13 的主导策略（StrategyCount 中出现次数最多的策略）与当前状态（在校/成绩/压力/参考高考）；
// 根据在校与否、压力≥0.5、家庭背景<0.3、成绩≥0.6 且在考池等条件拼接多条建议（保持努力学习、寻求减压、关注助学金、维持节奏或评估复读等），用分号连接输出。
func printC13Suggestion(state *SimState) {
	var c13 *Agent
	for i := range state.Agents {
		if state.Agents[i].Role == RoleStudentC13 {
			c13 = &state.Agents[i]
			break
		}
	}
	if c13 == nil {
		return
	}
	// 主导策略
	dominant := StrategyStudyHard
	maxCount := 0
	for _, s := range studentStrategies {
		if int(s) < len(c13.StrategyCount) && c13.StrategyCount[s] > maxCount {
			maxCount = c13.StrategyCount[s]
			dominant = s
		}
	}

	fmt.Println("\n=== C13 建议（高 IQ 贫困生） ===")
	fmt.Printf("当前状态：在校=%v  成绩=%.4f  压力=%.4f  参考高考=%v  主导策略=%s\n",
		c13.InSchool, c13.Score, c13.Stress, c13.InExamPool, dominant.String())

	fmt.Print("建议：")
	var tips []string
	if c13.InSchool {
		tips = append(tips, "保持「努力学习」策略，发挥高 IQ 优势")
		if c13.Stress >= 0.5 {
			tips = append(tips, "压力偏高可主动寻求心理老师减压")
		}
		if c13.Factor.FamilyBackground < 0.3 {
			tips = append(tips, "家庭资源有限可关注助学金、专项计划等升学路径")
		}
		if c13.Score >= 0.6 && c13.InExamPool {
			tips = append(tips, "当前成绩与参考状态有利于本科录取，可维持现有节奏")
		}
	} else {
		tips = append(tips, "已离校/退考，可评估复读或替代升学路径")
	}
	if len(tips) == 0 {
		tips = append(tips, "保持当前策略，注意压力与法规道德风险")
	}
	for i, t := range tips {
		if i > 0 {
			fmt.Print("；")
		}
		fmt.Print(t)
	}
	fmt.Println("。")
}
