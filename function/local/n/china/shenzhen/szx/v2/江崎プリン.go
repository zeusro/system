package v2

import (
	"fmt"
	"sync"
	"time"

	"github.com/zeusro/system/function/local/n/china/shenzhen/szx/model"
)

type Information struct {
	date    time.Time
	content string
}

type AlibabaGroup struct {
	N    int           //算法规模
	Cost time.Duration //总耗时
	model.Alipay
	model.Aliyun
}

func (a *AlibabaGroup) Actor(core string, inbox <-chan Information) {
	for msg := range inbox {
		fmt.Printf("[%v]Actor %s received[%d]: %s\n", msg.date, core, a.N, msg.content)
		a.N++
	}
}

// EatBean 如果把问题转换为一个整体（阿里云和支付宝同属阿里巴巴集团的资产），那么问题就可以简化为一个简单的生产者消费者模型
func (ali *AlibabaGroup) EatBean(beans []model.Bean) map[time.Time]model.Point {
	var m map[time.Time]model.Point = make(map[time.Time]model.Point)
	now := time.Now()
	start := now
	// fmt.Println(start)
	n := len(beans)
	var wg sync.WaitGroup
	memory := make(chan Information, 1) //限定为1，强转为同步队列结构
	wg.Add(1)
	go func() {
		defer wg.Done()
		ali.Actor("1A84", memory)
	}()
	a := model.RandonPoint()
	m[now] = a
	memory <- Information{content: "立ち上がれ、江崎プリン！"}
	memory <- Information{date: start, content: fmt.Sprintf("(%f,%f)", a.X, a.Y)}
	cache := make(map[float64]float64)
	for i := 0; i < n-1; i++ {
		b := model.RandonPoint()
		notIn := true
		for notIn {
			if _, contains := cache[b.X]; contains {
				b = model.RandonPoint()
				continue
			} else {
				cache[b.X] = b.Y
				notIn = false
				break
			}
		}
		line := model.NewLine(a, b)
		now = now.Add(line.Distance())
		m[now] = b
		memory <- Information{date: now, content: fmt.Sprintf("(%f,%f)", b.X, b.Y)}
	}
	// fmt.Println(now)
	// fmt.Printf("cost: %v \n", ali.GetCost())
	ali.Cost = now.Sub(start)
	memory <- Information{date: time.Now(), content: fmt.Sprintf("cost: %v", ali.GetCost())}
	close(memory)
	wg.Wait()
	return m
}

func (ali *AlibabaGroup) GetCost() time.Duration {
	return ali.Cost
}
