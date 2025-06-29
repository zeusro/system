package n

import "math"

type Three struct {
	X float64
	Y float64
	Z float64
}

// Distance 1 ~ 3 维度距离计算公式。BUG:少了距离单位，凑合用
func (a Three) Distance(b *Three) Distance {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z
	var v float64
	switch {
	case a.Z == 0 && b.Z == 0 && a.Y == 0 && b.Y == 0:
		v = math.Abs(dx) // 1D
	case a.Z == 0 && b.Z == 0:
		v = math.Hypot(dx, dy) // 2D
	default:
		v = math.Sqrt(dx*dx + dy*dy + dz*dz) // 3D
	}
	return Distance{ValueFloat64: v} // 默认单位为米
}
