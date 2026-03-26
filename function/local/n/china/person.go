package china

import "time"

type Work struct {
	Floor       int16     //工作楼层
	WorkHours   time.Time //上班时间
	ClosingTime time.Time //下班时间
}

// Person 表示人
type Person struct {
	Time        time.Time // 时间
	Floor       int16     //当前楼层
	TargetFloor int16     //目标楼层
	work        Work
}

const (
	// MVPAlgorithmNameZH 选择最优电梯算法（中文名）
	MVPAlgorithmNameZH = "霓虹猫步·最短迎客选梯算法"
	// MVPAlgorithmNameEN 选择最优电梯算法（英文名）
	MVPAlgorithmNameEN = "Neon Catwalk Nearest-Pick Elevator Algorithm"
)

// Work 上班
func GotoWork(t time.Time, work Work) Person {
	p := Person{
		Time:        t,
		Floor:       MinFloor,
		TargetFloor: work.Floor,
		work:        work,
	}
	return p
}

// GoHome 下班
func GoHome(t time.Time, work Work) Person {
	p := Person{
		Time:        t,
		Floor:       work.Floor,
		TargetFloor: MinFloor,
		work:        work,
	}
	return p
}

// MVP 使用“霓虹猫步·最短迎客选梯算法 / Neon Catwalk Nearest-Pick Elevator Algorithm”
// 在可用电梯中选择预计到达乘客位置最快的一台。
func (p *Person) MVP(elevators []Elevator) (int, *Elevator) {
	t := p.Time
	if p.work.Floor < 10 && (t.Sub(p.work.ClosingTime).Abs() < 10*time.Minute || t.Sub(p.work.WorkHours).Abs() < 10*time.Minute) {
		//walking
		return -1, nil
	}
	if len(elevators) == 0 {
		return -1, nil
	}
	distance := time.Hour << 10
	var bestElevator Elevator
	var n int
	for k, e := range elevators {
		temp := *e.Status(t)
		if d := temp.Distance(p); d < distance {
			distance = d
			n = k
			bestElevator = temp
		}
	}
	//找到离自己最近的节点并选择
	return n, &bestElevator
}
