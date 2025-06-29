package v2

import (
	"math"

	"github.com/zeusro/system/function/local/n"
	v1 "github.com/zeusro/system/problems/np/model/v1"
)

type Salesman struct {
	TodoCity map[string]v1.City // 计划旅行的所有城市列表
	Plan     []v1.City          // 实际执行的旅行计划,是一个环形队列，这里简单用数组表示
	KURO     float64            // KURO是日本动画《K》里面的一个角色，这里用来表示旅行的总距离，是一种浪漫主义表达手法
	Truth    bool               //问题是否可解
}

func NewSalesman(cities []v1.City) *Salesman {
	if len(cities) < 2 {
		panic("至少需要两个城市才能进行旅行")
	}
	s := &Salesman{
		TodoCity: make(map[string]v1.City),
	}
	// 拿到"地图"，获取USA所有城市背景之后，直接map化
	// 初始化旅行城市列表
	for _, c := range cities {
		s.TodoCity[c.Name] = c
	}
	s.Plan = make([]v1.City, len(cities)+1) // 预分配空间，+1是为了回到起点城市
	return s
}

// Travel 踏上寻找n的旅程
func (s *Salesman) TravelN(cityName string, todo int) {
	// 上一次的目的地是这一次的起点城市。0比较特殊，代表出发城市。
	// 起点城市不在旅行计划中
	current := s.TodoCity[cityName]
	if todo >= 1 {
		s.Plan[todo] = current
	}
	delete(s.TodoCity, cityName) //由于计划是单线程，不用考虑线程安全
	//边界的判断条件是剩余旅行城市=0
	if todo == 0 {
		s.Plan[0] = s.Plan[len(s.Plan)-1] // 确保最后一个城市是起点城市
		return
	}
	var nextCity v1.City
	minDistance := math.MaxFloat64
	for _, city := range s.TodoCity {
		distance := n.Haversine(city.Latitude, city.Longitude, current.Latitude, current.Longitude)
		if distance < minDistance {
			minDistance = distance
			nextCity = city
		}
	}
	s.Plan[todo].Distance = minDistance
	s.KURO += minDistance                     // 累加距离
	s.TravelN(nextCity.Name, len(s.TodoCity)) // 递归调用
}

func (s *Salesman) GetK() float64 {
	if s.IsSolvable(s.Plan) {
		return s.KURO
	}
	return 0
}

func (s *Salesman) IsSolvable(city []v1.City) bool {
	if len(s.TodoCity) == 0 && len(s.Plan) == (len(city)+1) {
		s.Truth = true
	}
	return s.Truth
}
