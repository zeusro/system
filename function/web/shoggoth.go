package web

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// YoungAndBeautiful 年轻又漂亮
type YoungAndBeautiful struct {
	Name           string  //敢问妹子芳名
	BaseEfficiency float64 // 基础追求效率
	DecayRate      float64 // 妹子对我的兴趣衰减率
	TimeSpent      float64 // 追妹子已花费时间(小时)
}

// ED 随着时间的增加，对妹子越来越不感兴趣，打算出家当和尚掩盖下面问题
// 计算当前学习效率
// CurrentEfficiency: 随着学习时间的增加，效率会逐渐下降，趋近于 0。
func (s *YoungAndBeautiful) ED() float64 {
	return s.BaseEfficiency * math.Exp(-s.DecayRate*s.TimeSpent)
}

// 计算学习效果(效率*时间)
func (s *YoungAndBeautiful) LearningEffect(time float64) float64 {
	return s.ED() * time
}

// EDOptimizer ED优化器
// 梯度下降优化器
type EDOptimizer struct {
	Subjects     []*YoungAndBeautiful
	TotalTime    float64
	LearningRate float64
	Iterations   int
}

// 优化时间分配
// most valuable policy
func (g *EDOptimizer) MostValuablePolicy() []float64 {
	// 初始化随机时间分配
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	timeAlloc := make([]float64, len(g.Subjects))
	sum := 0.0
	for i := range timeAlloc {
		timeAlloc[i] = r.Float64()
		sum += timeAlloc[i]
	}
	// 归一化
	for i := range timeAlloc {
		timeAlloc[i] = timeAlloc[i] / sum * g.TotalTime
	}

	for iter := 0; iter < g.Iterations; iter++ {
		// 计算当前总学习效果
		currentEffect := g.totalLearningEffect(timeAlloc)

		// 计算梯度
		gradients := make([]float64, len(g.Subjects))
		epsilon := 0.01

		for i := range g.Subjects {
			// 创建临时时间分配
			tempAlloc := make([]float64, len(timeAlloc))
			copy(tempAlloc, timeAlloc)
			tempAlloc[i] += epsilon

			// 归一化临时分配
			tempSum := 0.0
			for _, t := range tempAlloc {
				tempSum += t
			}
			for j := range tempAlloc {
				tempAlloc[j] = tempAlloc[j] / tempSum * g.TotalTime
			}

			// 计算梯度
			tempEffect := g.totalLearningEffect(tempAlloc)
			gradients[i] = (tempEffect - currentEffect) / epsilon
		}

		// 更新时间分配
		for i := range timeAlloc {
			timeAlloc[i] += g.LearningRate * gradients[i]
			timeAlloc[i] = math.Max(timeAlloc[i], 0) // 确保非负
		}

		// 重新归一化
		sum := 0.0
		for _, t := range timeAlloc {
			sum += t
		}
		for i := range timeAlloc {
			timeAlloc[i] = timeAlloc[i] / sum * g.TotalTime
		}
	}

	return timeAlloc
}

// 计算总学习效果
func (g *EDOptimizer) totalLearningEffect(timeAlloc []float64) float64 {
	total := 0.0
	for i, subj := range g.Subjects {
		total += subj.LearningEffect(timeAlloc[i])
	}
	return total
}

type Shoggoth struct {
	Dick      struct{} // 代表他的粉丝
	TotalDays int      // 模拟天数
	DailyTime float64  // 每天学习时间(小时)
}

// DynamicFindGirlfriends 根据效果动态（时间）调整泡妞策略
func (v Shoggoth) DynamicFindGirlfriends(subjects []*YoungAndBeautiful) {
	totalDays := v.TotalDays
	dailyTime := v.DailyTime

	for day := 1; day <= totalDays; day++ {
		fmt.Printf("\nDAY %d\n", day)

		// 创建优化器(每天重新优化)
		optimizer := EDOptimizer{
			Subjects:     subjects,
			TotalTime:    dailyTime,
			LearningRate: 0.1,
			Iterations:   50,
		}

		// 获取当天最优分配
		todayAlloc := optimizer.MostValuablePolicy()

		// 应用时间分配(更新已花费时间)
		for i, subj := range subjects {
			subj.TimeSpent += todayAlloc[i]
		}
		gotED := true
		// 打印当天计划
		fmt.Println("今日推荐学习时间:")
		for i, subj := range subjects {
			ed := subj.ED()
			fmt.Printf("%s: %.2f小时 (累计已学: %.2f小时, 当前效率: %.2f)\n",
				subj.Name, todayAlloc[i], subj.TimeSpent, ed)
			//在 Go 中，float64 类型的数值计算是不精确的。
			//即使一个看起来为 0.0 的值，实际可能是 1e-10、-1e-12 等极小数，不等于 0.00 精确值。
			if ed > 1e-6 {
				gotED = false
			}
		}
		if gotED {
			fmt.Println("所有妹子都不感兴趣了，打算出家当和尚。")
			return
		}

		// 模拟效率变化(随机波动)
		for _, subj := range subjects {
			// 添加一些随机波动
			subj.BaseEfficiency *= 0.95 + 0.1*rand.Float64()
		}
	}
}
