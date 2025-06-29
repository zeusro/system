package n

// Location 位置
type Location interface {
	//Where are you?
}

type Distance struct {
	DistanceUnit //单位 enum
	//数值 decimal
	ValueFloat64 float64
	//TODO 按需补充
}

// 基础类型定义：DistanceUnit 为 int 类型
type DistanceUnit int

const (
	// 常见单位，按常用缩写定义
	Millimeter DistanceUnit = iota
	Centimeter
	Meter
	Kilometer
	Inch
	Foot
	Yard
	Mile
	NauticalMile
)

func (du DistanceUnit) String() string {
	switch du {
	case Millimeter:
		return "mm"
	case Centimeter:
		return "cm"
	case Meter:
		return "m"
	case Kilometer:
		return "km"
	case Inch:
		return "in"
	case Foot:
		return "ft"
	case Yard:
		return "yd"
	case Mile:
		return "mi"
	case NauticalMile:
		return "nmi"
	default:
		return "unknown"
	}
}
