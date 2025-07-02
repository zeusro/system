package v1

const (
	//跑道600m
	distance int64 = 600
	//时间限制3600s
	limit int64 = 3600
)

// Zeusro 我的滑板车
func Zeusro(limit int64) (timings []int64) {
	timings = append(timings, 0)
	var v1 int64 = 1
	cycle := distance / v1
	var time int64 = cycle
	//速度*时间=路程
	for time = cycle; time < limit; time += cycle {
		timings = append(timings, time)
	}
	return timings
}

// YoungAndBeautiful 妹子慢慢走
func YoungAndBeautiful(limit int64) (timings []int64) {
	timings = append(timings, 0)
	var v2 int64 = 2
	cycle := distance / v2
	var time int64 = cycle
	for time = cycle; time < limit; time += cycle {
		timings = append(timings, time)
	}
	return timings
}

// RichGrandma 富婆在跑步
func RichGrandma(limit int64) (timings []int64) {
	timings = append(timings, 0)
	var v3 int64 = 3
	cycle := distance / v3
	var time int64 = cycle
	for time = cycle; time < limit; time += cycle {
		timings = append(timings, time)
	}
	return timings
}
