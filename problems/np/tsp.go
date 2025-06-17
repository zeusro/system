package np

import (
	"math"
)

type Salesman struct {
	TodoCity map[string]City // 计划旅行的所有城市列表
	Plan     []City          // 实际执行的旅行计划,是一个环形队列，这里简单用数组表示
}

func NewSalesman() *Salesman {
	s := &Salesman{
		TodoCity: make(map[string]City),
		Plan:     make([]City, 0),
	}
	// 拿到"地图"，获取USA所有城市背景之后，直接map化
	// 初始化旅行城市列表
	for _, c := range usCities {
		s.TodoCity[c.Name] = c
	}
	return s
}

// Travel 踏上旅程，寻找真我
func (s *Salesman) Travel(current City, plan []City) []City {
	// 删除起点城市
	// /上一次的目的地是这一次的起点城市
	delete(s.TodoCity, current.Name) //由于计划是单线程，不用考虑线程安全
	n := len(s.TodoCity)
	if n == 1 {
		s.Plan = append(s.Plan, current)
	}
	//边界的判断条件是剩余旅行城市=0
	if n == 0 {
		s.Plan = append(s.Plan, s.Plan[0]) // 回到起点，形成环形
		return s.Plan
	}
	var nextCity City
	minDistance := math.MaxFloat64
	// todo:如果“n”的范围很大，这里可以用经纬度上下界,以current作为中心点，限定计算网格大小，从而方便更快地遍历穷举
	// 用SQL表示就是 select citys from USA where c.Latitude between 24.5 and 49.4 and c.Longitude between -124.8 and -66.9
	// 不过这种传统关系型数据库，查询效率不符合我的要求
	for _, city := range s.TodoCity { //fixme：当前的数组集合类型是有缺陷的，不能一次性全部取出，导致了O(n)的算法复杂度，实际上应该是O(1)然后并发算出最小距离城市
		distance := haversine(city.Coordinates.Latitude, city.Coordinates.Longitude, current.Coordinates.Latitude, current.Coordinates.Longitude)
		if distance < minDistance {
			minDistance = distance
			nextCity = city
		}
	}
	s.Plan = append(s.Plan, nextCity)

	return s.Travel(nextCity, plan) // 递归调用
}

// haversine 📌 Haversine 公式：计算地球上两点的距离
// 传入两点的经纬度，返回两点之间的距离（单位：公里）
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // 地球半径（单位：公里）

	dLat := degreesToRadians(lat2 - lat1)
	dLon := degreesToRadians(lon2 - lon1)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(degreesToRadians(lat1))*math.Cos(degreesToRadians(lat2))*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func degreesToRadians(deg float64) float64 {
	return deg * math.Pi / 180
}
