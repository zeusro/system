package china

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// Tendency 定义电梯运行趋势的枚举类型
type Tendency int

const (
	StatusEmpty Tendency = iota // 0: 空载状态 (待机)
	StatusUp                    // 1: 上
	StatusDown                  // 2: 下
)

const MinFloor int16 = 1
const DefaultTopFloor int16 = 30

func (t Tendency) String() string {
	switch t {
	case StatusEmpty:
		return "空载待机"
	case StatusUp:
		return "向上运行"
	case StatusDown:
		return "向下运行"
	default:
		return "未知状态"
	}
}

// Elevator 表示电梯
type Elevator struct {
	Time         time.Time  // 时间
	CurrentFloor int16      // 当前楼层
	Tendency     Tendency   // 运行趋势
	Capacity     int16      // 载客量
	Top          int16      //楼层上限
	Data         []Elevator // 环形队列的底层数组，用于存储运行轨迹
	Front        int        // 环形队列的头指针
	Rear         int        // 环形队列的尾指针
	MaxQueueSize int        // 环形队列的最大容量（实际可用容量为 MaxQueueSize - 1）
	queueMu      *sync.RWMutex
}

// RandomElevator 生成电梯的随机运行状态
func RandomElevator(start time.Time, top int16, queueSize int) *Elevator {
	if top < MinFloor {
		top = DefaultTopFloor
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	currentFloor := int16(r.Intn(int(top-MinFloor+1)) + int(MinFloor))
	tendency := Tendency(r.Intn(3))
	e := NewElevator(start, currentFloor, tendency, queueSize)
	e.Top = top
	return e
}

func NewElevator(start time.Time, currentFloor int16, tendency Tendency, n int) *Elevator {
	e := Elevator{
		Time:         start,
		CurrentFloor: currentFloor,
		Tendency:     tendency,
		Capacity:     13,
		Top:          DefaultTopFloor,
		Data:         make([]Elevator, n),
		// data读写➕读写锁
		queueMu: &sync.RWMutex{},
	}
	e.initQueue(n)
	return &e
}

// InitQueue 初始化电梯内部的环形队列
func (e *Elevator) initQueue(size int) {
	e.MaxQueueSize = size + 1 // 为了区分空和满，实际申请 size+1 的空间
	e.Data = make([]Elevator, e.MaxQueueSize)
	e.Front = 0
	e.Rear = 0
}

// Enqueue 将电梯状态加入环形队列
func (e *Elevator) Enqueue(state Elevator) {
	e.queueMu.Lock()
	defer e.queueMu.Unlock()

	// 如果队列已满，则将头指针向后移动一位（覆盖最老的数据）
	if (e.Rear+1)%e.MaxQueueSize == e.Front {
		e.Front = (e.Front + 1) % e.MaxQueueSize
	}
	e.Data[e.Rear] = state
	e.Rear = (e.Rear + 1) % e.MaxQueueSize
}

// PrintHistory 打印环形队列中记录的电梯运行轨迹
func (e *Elevator) PrintHistory() {
	fmt.Println("=== 电梯运行轨迹记录 ===")
	if e.Front == e.Rear {
		fmt.Println("无记录")
		return
	}
	i := e.Front
	for i != e.Rear {
		state := e.Data[i]
		fmt.Printf("时间: %s | 楼层: %d | 状态: %s\n", state.Time.Format("15:04:05"), state.CurrentFloor, state.Tendency.String())
		i = (i + 1) % e.MaxQueueSize
	}
	fmt.Println("========================")
}

func (e *Elevator) Status(t time.Time) *Elevator {
	// 从环形队列中获取电梯当前的状态
	e.queueMu.RLock()
	defer e.queueMu.RUnlock()

	// 没有历史记录时返回当前状态
	if e.Front == e.Rear {
		snapshot := *e
		return &snapshot
	}

	i := e.Front
	var chosen Elevator
	var picked bool
	var fallback Elevator
	var hasFallback bool
	for i != e.Rear {
		state := e.Data[i]
		if !state.Time.IsZero() {
			if state.Time.Before(t) || state.Time.Equal(t) {
				if !picked || state.Time.After(chosen.Time) {
					chosen = state
					picked = true
				}
			} else if !hasFallback || state.Time.Before(fallback.Time) {
				// 如果没有 <= t 的状态，使用最早的未来状态兜底
				fallback = state
				hasFallback = true
			}
		}
		i = (i + 1) % e.MaxQueueSize
	}

	if picked {
		return &chosen
	}
	if hasFallback {
		return &fallback
	}
	snapshot := *e
	return &snapshot
}

func (e *Elevator) Distance(p *Person) time.Duration {
	if p.TargetFloor == p.Floor {
		return 0
	}

	var tendency Tendency
	if p.TargetFloor > p.Floor {
		tendency = StatusUp
	} else {
		tendency = StatusDown
	}

	top := e.Top
	if top < MinFloor {
		top = DefaultTopFloor
	}

	var floor int16
	if e.Tendency == StatusEmpty {
		floor = int16(math.Abs(float64(e.CurrentFloor - p.Floor)))
	} else if tendency == e.Tendency {
		//同向上行取最小值
		if tendency == StatusUp {
			if e.CurrentFloor <= p.Floor {
				floor = p.Floor - e.CurrentFloor
			} else {
				floor = (top - e.CurrentFloor) + (top - p.Floor)
			}
		} else {
			if e.CurrentFloor >= p.Floor {
				floor = e.CurrentFloor - p.Floor
			} else {
				floor = (e.CurrentFloor - MinFloor) + (p.Floor - MinFloor)
			}
		}
	} else {
		// 不同方向：需要先到当前方向边界再折返
		if e.Tendency == StatusUp {
			floor = (top - e.CurrentFloor) + (top - p.Floor)
		} else {
			floor = (e.CurrentFloor - MinFloor) + (p.Floor - MinFloor)
		}
	}

	return (time.Hour * time.Duration(floor))
}

// Run 模拟电梯从初始状态 (start) 运行到当前电梯 (e) 的目标状态
func (e *Elevator) Run(start Elevator) *Elevator {
	// 使用局部随机源，替换已过时的全局 rand.Seed 用法
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 从 start 状态开始模拟
	currentState := start

	// 确定运行方向
	if currentState.CurrentFloor < e.CurrentFloor {
		currentState.Tendency = StatusUp
	} else if currentState.CurrentFloor > e.CurrentFloor {
		currentState.Tendency = StatusDown
	} else {
		currentState.Tendency = StatusEmpty
	}

	fmt.Printf("开始运行：从 %d 层前往 %d 层...\n", currentState.CurrentFloor, e.CurrentFloor)

	// 记录初始状态
	e.Enqueue(currentState)

	// 逐层运行直到到达目标楼层
	for currentState.CurrentFloor != e.CurrentFloor {
		// 生成 1~5 秒的随机运行时间
		randomSeconds := r.Intn(5) + 1
		duration := time.Duration(randomSeconds) * time.Second

		// 模拟时间流逝
		currentState.Time = currentState.Time.Add(duration)
		time.Sleep(duration) // 如果只想看计算结果，可以注释掉这行 Sleep

		// 更新楼层
		if currentState.Tendency == StatusUp {
			currentState.CurrentFloor++
		} else if currentState.Tendency == StatusDown {
			currentState.CurrentFloor--
		}

		fmt.Printf("经过楼层: %d (耗时: %d秒, 当前时间: %s)\n", currentState.CurrentFloor, randomSeconds, currentState.Time.Format("15:04:05"))

		// 将每层的状态存入环形队列
		e.Enqueue(currentState)
	}

	// 到达目标状态后，设为待机
	currentState.Tendency = StatusEmpty
	e.Time = currentState.Time // 更新目标电梯的最终时间
	e.Enqueue(currentState)
	fmt.Println("电梯已到达目标楼层。")

	return e
}
