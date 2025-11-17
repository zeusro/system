package v3

import (
	"fmt"
	"time"

	"github.com/zeusro/system/function/local/n/china/shenzhen/szx/model"
)

func DoubleThought(n int) []NLine {
	m := make(map[int]model.Point, n)
	for i := 1; i < n; i++ {
		m[i] = model.RandonPoint()
	}
	p1 := model.RandonPoint()
	p2 := model.RandonPoint()
	//强制分配不同的起点
	for p1.Compare(p2) {
		p2 = model.RandonPoint()
	}
	bean1 := NewBeansWithFirstPoint(p1, m)
	bean1.Name = "aliyun"
	bean2 := NewBeansWithFirstPoint(p2, m)
	bean2.Name = "alipay"

	now := time.Now()
	journey1 := bean1.Thought(n, now)
	journey2 := bean2.Thought(n, now)
	//归一化合并，去掉多余的点。
	nLines := make(map[time.Time]NLine)
	nMap := NewNLineMap(0)
	for k, t1 := range journey1.NBeans {
		t2, ok := journey2.NBeans[k]
		if !ok {
			continue
		}
		//落后就要挨打
		if t1.After(t2) {
			delete(journey1.NBeans, k)
			line := NLine{t: t2, Line: k.Line, actorID: bean2.Name}
			nLines[t2] = line
			nMap.Add(t2, line)
			continue
		}
		if t2.After(t1) {
			delete(journey2.NBeans, k)
			line := NLine{t: t1, Line: k.Line, actorID: bean1.Name}
			nLines[t1] = line
			nMap = nMap.Add(t1, line)
			continue
		}
		//由于字典会自动去重，因此碰撞可以忽略
		fmt.Printf("%v:%v 两个吃豆人同时到达%v，发生碰撞\n", t1, t2, k.Line)
	}
	nMap.AddZero(bean1.FirstNL)
	nMap.AddZero(bean2.FirstNL)
	fmt.Printf("%v : n维线段总数：%d;吃豆人%v总数%v;吃豆人%v总数%v;\n", now,
		len(nMap.items)+2, bean1.Name, len(journey1.NBeans), bean2.Name, len(journey2.NBeans))
	lines := nMap.All(false)
	fmt.Printf("len(lines):%v\n", len(lines))
	for k, v := range lines {
		fmt.Println(v.String(k))
		// fmt.Printf("%v:%v\n", k, v.String(k))
	}
	fmt.Println(journey1.End())
	return lines
}
