package local

import "math"

// haversine 📌 Haversine 公式：计算地球上两点的距离
// 传入两点的经纬度，返回两点之间的距离（单位：公里）
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
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
