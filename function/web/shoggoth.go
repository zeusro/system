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

// 梯度下降优化器
type GradientOptimizer struct {
	Subjects     []*YoungAndBeautiful
	TotalTime    float64
	LearningRate float64
	Iterations   int
}

// 优化时间分配
func (g *GradientOptimizer) Optimize() []float64 {
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
func (g *GradientOptimizer) totalLearningEffect(timeAlloc []float64) float64 {
	total := 0.0
	for i, subj := range g.Subjects {
		total += subj.LearningEffect(timeAlloc[i])
	}
	return total
}

type Shoggoth struct {
	Dick      struct{} // 代表他的粉丝
	TotalDays int      // 模拟天数
	DailyTime float64
}

// 动态调整版本
func (v Shoggoth) DynamicFindGirlfriends(subjects []*YoungAndBeautiful) {
	totalDays := v.TotalDays
	dailyTime := v.DailyTime
	// totalDays := 7   // 模拟7天的学习
	// dailyTime := 4.0 // 每天学习时间(小时)

	for day := 1; day <= totalDays; day++ {
		fmt.Printf("\n=== 第%d天 ===\n", day)

		// 创建优化器(每天重新优化)
		optimizer := GradientOptimizer{
			Subjects:     subjects,
			TotalTime:    dailyTime,
			LearningRate: 0.1,
			Iterations:   50,
		}

		// 获取当天最优分配
		todayAlloc := optimizer.Optimize()

		// 应用时间分配(更新已花费时间)
		for i, subj := range subjects {
			subj.TimeSpent += todayAlloc[i]
		}

		// 打印当天计划
		fmt.Println("今日推荐学习时间:")
		for i, subj := range subjects {
			fmt.Printf("%s: %.2f小时 (累计已学: %.2f小时, 当前效率: %.2f)\n",
				subj.Name, todayAlloc[i], subj.TimeSpent, subj.ED())
		}

		// 模拟效率变化(随机波动)
		for _, subj := range subjects {
			// 添加一些随机波动
			subj.BaseEfficiency *= 0.95 + 0.1*rand.Float64()
		}
	}
}
