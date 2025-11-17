package v3

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeusro/system/function/local/n/china/shenzhen/szx/model"
)

// go test -run TestDoubleThought -v
func TestDoubleThought(t *testing.T) {
	lines := DoubleThought(50)
	t.Logf("len(lines)：%v", len(lines))
	// for _, line := range lines {
	// 	t.Log(line.String())
	// }
	// t.Log(lines)
}

func TestThought(t *testing.T) {
	n := 50
	m := make(map[int]model.Point, n)
	for i := 1; i < n; i++ {
		m[i] = model.RandonPoint()
	}
	p1 := model.RandonPoint()
	bean1 := NewBeansWithFirstPoint(p1, m)
	now := time.Now()
	m1 := bean1.Thought(n, now)
	for k, v := range m1.Lines {
		// t.Logf("%v:%v\n", v, k)
		now = now.Add(m1.Lines[k].Distance())
		fmt.Printf("%v:%v\n", now, v)
	}
}
