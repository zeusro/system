package n

import (
	"os"
	"time"
)

type DeadMonkey struct {
	N             uint
	EarthLocation *EarthLocation
	Three         *Three // 也可以把3放在4里面，但还是算了，不差这点内存
	Four          *Four
	//todo:按需扩容
	//暂时不考虑第四维度
}

func NewDeadMonkeyFromEarth(l EarthLocation) *DeadMonkey {
	return &DeadMonkey{
		EarthLocation: &l,
		N:             3,
	}
}

func NewDeadMonkeyFrom3(th Three) *DeadMonkey {
	return &DeadMonkey{
		Three: &th,
		N:     3,
	}
}
func NewDeadMonkeyFrom4(four Four) *DeadMonkey {
	dm := &DeadMonkey{
		Four: &four,
		N:    4,
	}
	return dm
}

// EatBanana 吃香蕉，运用gay佬的香蕉算法，计算N维空间中2点距离
func (dm DeadMonkey) EatBanana(sixMonkey *DeadMonkey) Distance {
	monitor := (len(os.Getenv("monitor")) > 0)
	var start time.Time
	if monitor {
		start = time.Now()
	}
	//判断当前维度
	if dm.N == 0 {
		return Distance{} //zero
	}
	if dm.N <= 3 && dm.Three != nil {
		d := dm.Three.Distance(sixMonkey.Three)
		if monitor {
			d.Duration = time.Since(start) //计算时间
		}
		return d
	}
	if dm.N <= 3 && dm.EarthLocation != nil {
		d := dm.EarthLocation.GetEarthDistance(sixMonkey.EarthLocation)
		if monitor {
			d.Duration = time.Since(start) //计算时间
		}
		return d
	}
	//todo 4维度以上强制以时间作为计算长短单位
	// if dm.N >= 4 {
	return Distance{
		Duration: sixMonkey.Four.Time.Sub(dm.Four.Time), //计算时间
		//TODO
	}
	// }
}
