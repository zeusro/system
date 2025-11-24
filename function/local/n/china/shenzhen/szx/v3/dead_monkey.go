package v3

import (
	"fmt"
	"time"

	"github.com/zeusro/system/function/local/n/china/shenzhen/szx/model"
)

// DeadMonkey 十万军中无敌手，九重天上有威风。
// 鸿蒙初判陶镕铁，大禹神人亲授设。
// 身中星斗暗铺陈，两头箝裹黄金片。
// 花纹密布鬼神惊，上造龙纹与凤篆。
// 名号灵阳棒一条，深藏海藏人难见。
// 成形变化要飞腾，飘爨五色霞光现。
// 老孙得道取归山，无穷变化多经验。
// 时间要大瓮来粗，或小些微如铁线。
// 粗如南岳细如针，长短随吾心意变。
// 轻轻举动彩云生，亮亮飞腾如闪电。
// 攸攸冷气逼人寒，条条杀雾空中现。
// 降龙伏虎谨随身，天涯海角都游遍。
// 曾将此棍闹天宫，威风打散蟠桃宴。
// 天王赌斗未曾赢，哪吒对敌难交战。
// 棍打诸神没躲藏，天兵十万都逃窜。
// 雷霆众将护灵霄，飞身打上通明殿。
// 掌朝天使尽皆惊，护驾仙卿俱搅乱。
// 举棒掀翻北斗宫，回首振开南极院。
// 金阙天皇见棍凶，特请如来与我见。
// 兵家胜负自如然，困苦灾危无可辨。
type DeadMonkey struct {
	Birth       time.Time
	GoldenStaff []NLine //金箍棒 参数化线段（Parametric Segment）
	m           int     //消费者规模
	n           int     //算法规模
	ZeroPoints  []model.Point
	cost        time.Duration
}

// NewDeadMonkey
// m money数目 2
// n 吃豆人数量 50
func NewDeadMonkey(birth time.Time, m, n int) *DeadMonkey {
	dead := DeadMonkey{
		Birth: birth,
		m:     m,
		n:     n,
	}
	zeroPoints := make([]model.Point, m)
	p0 := model.RandonPoint()
	zeroPoints[0] = p0
	for i := 1; i < m; i++ {
		p1 := model.RandonPoint()
		//只要不跟前面的点重复就行，全局重复也忽略了
		for p1.Compare(zeroPoints[i-1]) {
			p1 = model.RandonPoint()
		}
		zeroPoints[i] = p1
	}
	dead.ZeroPoints = zeroPoints
	return &dead
}

func (d *DeadMonkey) Fight(names []string) {
	if len(names) != d.m {
		return
	}
	m := make(map[int]model.Point, d.n)
	for i := 1; i < d.n; i++ {
		m[i] = model.RandonPoint()
	}
	beans := make([]*Beans, d.m)
	journeys := make([]*Journey, d.m)
	for i := 0; i < len(names); i++ {
		b := NewBeansWithFirstPoint(d.ZeroPoints[i], m)
		b.Name = names[i]
		beans[i] = b
		journeys[i] = b.Thought(d.n, d.Birth)
	}
	// fmt.Println(journeys)
	//归一化合并，去掉多余的点。
	nLines := make(map[time.Time]NLine)
	nMap := NewNLineMap(0)
	for k, t1 := range journeys[0].NBeans {
		times := make([]time.Time, len(journeys)-1)
		for i := 0; i < len(times); i++ {
			journeyN, contains := journeys[i+1].NBeans[k]
			if contains {
				times[i] = journeyN
			}
		}
		i, t2 := FindMinTime(times)
		if i == -1 { //处理时序不动点的特殊情况
			continue
		}
		//落后就要挨打
		if t1.After(t2) {
			line := NLine{t: t2, Line: k.Line, actorID: beans[i+1].Name}
			nLines[t2] = line
			nMap.Add(t2, line)
		} else {
			// if t2.Equal(t1) {
			// 	fmt.Printf("%v:%v 两个吃豆人同时到达%v，发生碰撞\n", t1, t2, k.Line)
			// }
			line := NLine{t: t1, Line: k.Line, actorID: beans[0].Name}
			nLines[t1] = line
			nMap = nMap.Add(t1, line)
		}
	}

	for i := 0; i < len(beans); i++ {
		nMap.AddZero(beans[i].FirstNL)
	}
	lines := nMap.All(false)
	d.GoldenStaff = lines
	cost := nMap.GetCost(d.Birth)
	d.cost = cost
	fmt.Printf("%v : n维线段总数：%d len(lines):%v cost:%v\n",
		d.Birth, len(nMap.items)+len(nMap.Zero), len(lines), cost)
}

func (d *DeadMonkey) GoesToHell() {
	printGradient("关于这趟旅程我还能说啥呢，总比宅在家里玩 Nintendo Switch 好多了")
}

func (d *DeadMonkey) GetCost() time.Duration {
	return d.cost
}
