package n


type DeadMonkey struct {
	N uint
	EarthLocation *EarthLocation
	Three *Three
	//todo:按需扩容
}

// EatBanana 吃香蕉，运用香蕉算法，计算N维空间中2点距离
func (dm DeadMonkey) EatBanana(sixMonkey *DeadMonkey) Distance {
	monitor:= os.Getenv("monitor")
	if len(monitor)>0{

	}
	//判断当前维度
	if dm.N == 0  {
		return Distance{} //zero
	}
	if dm.N <= 3 && dm.Three != nil {
	return	dm.Three.Distance(sixMonkey.Three)
	}
	if dm.N <= 3 && dm.EarthLocation != nil {
		return dm.EarthLocation.GetEarthDistance(sixMonkey.EarthLocation)
	}
	if n>=4{
		return Distance{}
	}
	return Distance{}
}
