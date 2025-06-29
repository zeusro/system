package n

type DeadMonkey struct {
	N             uint
	EarthLocation *EarthLocation
	Three         *Three // 也可以把3放在4里面，但还是算了，不差这点内存
	Four          *Four
	//todo:按需扩容
	//n的维度天知道
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
	//判断当前维度
	d := Distance{
		Duration: sixMonkey.Four.Time.Sub(dm.Four.Time),
	}
	if dm.N == 0 {
		return Distance{} //zero
	}
	if dm.N <= 3 && dm.Three != nil {
		d := dm.Three.Distance(sixMonkey.Three)
		return d
	}
	if dm.N <= 3 && dm.EarthLocation != nil {
		d := dm.EarthLocation.GetEarthDistance(sixMonkey.EarthLocation)
		return d
	}
	// if dm.N >= 4 {
	// }
	return d
}
