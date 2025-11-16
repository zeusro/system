package v1

import (
	"maps"
	"sync"
	"time"

	"github.com/zeusro/system/function/local/n/china/shenzhen/szx/model"
)

// Beans 结构：内置并发字典 key->Point
type Beans struct {
	mu    sync.RWMutex
	items map[int]model.Point
}

// NewBeans 根据输入 keys 和对应的 Point 值来初始化 Beans。
// 例如传入一个 map[int]Point 或是 slice 来构建。
func NewBeans(initial map[int]model.Point) *Beans {
	b := &Beans{
		items: make(map[int]model.Point, len(initial)),
	}
	maps.Copy(b.items, initial)
	return b
}

func NewBeansWithFirstPoint(first model.Point, initial map[int]model.Point) *Beans {
	initial[0] = first
	// initial[1] = second
	return NewBeans(initial)
}

// GetAndRemove 从字典中读取 key 对应的 Point，读取后删除该 key。
// 返回 value 和一个 bool 表示是否存在。
func (b *Beans) GetAndRemove(key int) (model.Point, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	v, ok := b.items[key]
	if !ok {
		return model.Point{}, false
	}
	delete(b.items, key)
	return v, true
}

// Exists 检查 key 是否存在（不删除）。
func (b *Beans) Exists(key int) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	_, ok := b.items[key]
	return ok
}

// Len 返回当前字典中剩余的元素数。
func (b *Beans) Len() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return len(b.items)
}

// Thought 以简单贪心算法进行吃豆人路径规划，有bug，但是我不修
func (beans *Beans) Thought(n int, date time.Time) map[model.Bean]time.Time {
	//简化问题，把随机初始点作为吃豆人起点
	a := make([]model.Point, 1)
	first, _ := beans.GetAndRemove(0)
	a[0] = first
	c := NewAlibabaCompany()
	for i := 1; i < n; i++ {
		p, contains := beans.GetAndRemove(i)
		if contains {
			line := model.NewLine(a[len(a)-1], p)
			a = append(a, p)
			c.Lines = append(c.Lines, line)
		}
	}
	aliBeans := make([]model.Bean, len(c.Lines))
	for k, item := range c.Lines {
		aliBeans[k] = model.Bean{Line: item}
	}
	return c.GetBeans(date, aliBeans)
}
