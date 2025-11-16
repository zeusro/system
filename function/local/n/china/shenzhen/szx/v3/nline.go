package v3

import (
	"fmt"
	"sort"
	"time"

	"github.com/zeusro/system/function/local/n/china/shenzhen/szx/model"
)

// NLine 基于时间的n维线段
type NLine struct {
	t       time.Time
	actorID string
	model.Line
}

func (l NLine) String() string {
	return fmt.Sprintf("%v %s:(%f,%f)-(%f,%f)", l.t, l.actorID, l.A.X, l.A.Y, l.B.X, l.B.Y)
}

type NLineMap struct {
	Zero  []NLine //起点是特殊的N维空间零点,存在N个时间相同的值
	items map[time.Time]NLine
}

func NewNLineMap(n int) *NLineMap {
	m := &NLineMap{
		Zero: make([]NLine, 0),
	}
	if n == 0 {
		m.items = make(map[time.Time]NLine)
	} else {
		m.items = make(map[time.Time]NLine, n)
	}
	return m
}

func (m *NLineMap) AddZero(nline NLine) *NLineMap {
	m.Zero = append(m.Zero, nline)
	return m
}

func (m *NLineMap) Add(t time.Time, nline NLine) *NLineMap {
	m.items[t] = nline
	return m
}

func (m *NLineMap) SortKeys() []time.Time {
	keys := make([]time.Time, len(m.items))
	i := 0
	for k, _ := range m.items {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j]) // 升序
	})
	return keys
}

func (m *NLineMap) All(print bool) []NLine {
	lenZero := len(m.Zero)
	nlines := make([]NLine, len(m.items)+lenZero)
	copy(nlines, m.Zero)
	keys := m.SortKeys()
	for i, k := range keys {
		nlines[i+lenZero] = m.items[k]
		// value, contains := m.items[k]
		// if contains {
		// 	nlines[i+lenZero] = value
		// } else {
		// 	fmt.Printf("error: not contains %v:%v\n", i+lenZero, k)
		// }
	}
	if print {
		for _, line := range nlines {
			fmt.Printf("%v:%v：(%f,%f)-(%f,%f)\n", line.t, line.actorID, line.A.X, line.A.Y, line.B.X, line.B.Y)
		}
	}
	return nlines
}
