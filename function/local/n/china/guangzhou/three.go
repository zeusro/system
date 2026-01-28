package guangzhou

import (
	"fmt"
	"math/rand"
	"time"
)

// Coke 算法，纪念一下广州金闪闪
// min 批发价
// max 最高零售价
// near  附近商户数量
// deviation  允许的合理偏差值
func Coke(min, max float64, near int, deviation float64) {
	//简化问题，以500ml可乐为例
	arround := RandomUniqueArray(near, min, max)  //调查获得附近500米内小商户，假设有15家
	juan := MinInArray(arround)                   //卷王
	limit := MaxInArray(arround, juan, deviation) //把价格差控制在5毛以内
	finalPrice := limit + deviation
	fmt.Printf("附近商户卖%v\n", arround)
	fmt.Println("你们不让我卖，老子就非要卖！你有意见你来卖！")
	fmt.Printf("卷王卖%v。我卖%v。\n", juan, finalPrice)
	fmt.Println("更少的工作时间，更高的利润率，让卷王不服不行。")
}

// RandomRange 生成在给定 min 和 max 范围内的一个新的随机上下界（保留到0.1精度）
func RandomRange(min, max float64) (float64, float64) {
	if min >= max {
		panic("min must be less than max")
	}
	rand.Seed(time.Now().UnixNano())
	scale := 10.0
	a := float64(int((min+rand.Float64()*(max-min))*scale)) / scale
	b := float64(int((min+rand.Float64()*(max-min))*scale)) / scale
	if a > b {
		a, b = b, a
	}
	return a, b
}

// RandomUniqueArray 生成一个长度为 n 的不重复浮点数组，数值在 RandomRange(min, max) 范围内（保留到0.01精度）
func RandomUniqueArray(n int, min, max float64) []float64 {
	if n <= 0 {
		return []float64{}
	}

	low, high := RandomRange(min, max)
	if high-low < 0.01 {
		panic("range too narrow to generate unique values")
	}

	precision := 100.0 // 对应 0.01 精度
	maxUnique := int((high - low) * precision)
	if n > maxUnique {
		// panic("requesting more unique numbers than the range can support at 0.01 precision")
	}

	rand.Seed(time.Now().UnixNano())
	result := make([]float64, 0, n)
	seen := make(map[float64]struct{})

	for len(result) < n {
		val := low + rand.Float64()*(high-low)
		rounded := float64(int(val*precision)) / precision // 限制精度到 0.01
		if _, exists := seen[rounded]; !exists {
			seen[rounded] = struct{}{}
			result = append(result, rounded)
		}
	}
	return result
}

// MaxInArray 返回满足条件（不超过 min+deviation）的最大值
func MaxInArray(arr []float64, min, deviation float64) float64 {
	if len(arr) == 0 {
		panic("array is empty")
	}
	limit := min + deviation
	found := false
	max := arr[0]
	for _, v := range arr {
		if v <= limit {
			if !found || v > max {
				max = v
				found = true
			}
		}
	}
	if !found {
		// panic("no values within the allowed range")
	}
	return max
}

// MinInArray 返回 float64 数组中的最小值
func MinInArray(arr []float64) float64 {
	if len(arr) == 0 {
		panic("array is empty")
	}
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
