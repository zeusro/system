package n

import (
	"fmt"
	"testing"
	"time"
)

func TestEatBanana(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t0, _ := time.ParseInLocation("2006-01-02 15:04:05", "2006-01-02 15:04:05", loc)

	dead := NewDeadMonkeyFrom3(Three{X: 0})
	dead.Four = &Four{Time: t0}
	six := NewDeadMonkeyFrom3(Three{X: 984729})
	distance := dead.EatBanana(six)
	fmt.Printf("distance struct: %#v\n", distance)

	dead.N = 4
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2025-06-29 14:30:00", loc)
	four := Four{Time: t1}
	six = NewDeadMonkeyFrom4(four)
	distance = dead.EatBanana(six)
	fmt.Printf("4D distance: %#v\n", distance)

	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-06-29 14:30:00", loc)

	four = Four{Time: t2}
	six = NewDeadMonkeyFrom4(four)
	distance = dead.EatBanana(six)
	fmt.Printf("4D distance: %#v\n", distance)
}
