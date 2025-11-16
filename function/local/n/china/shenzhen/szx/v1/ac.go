package v1

import (
	"time"

	"github.com/zeusro/system/function/local/n/china/shenzhen/szx/model"
)

type AlibabaCompany struct {
	model.Alipay
	model.Aliyun
	Lines []model.Line
}

func NewAlibabaCompany() *AlibabaCompany {
	c := &AlibabaCompany{
		Lines: make([]model.Line, 0),
	}
	return c
}

func (ali *AlibabaCompany) EatBeans(beans []model.Bean) map[time.Time]model.Bean {
	m := make(map[time.Time]model.Bean)
	now := time.Now()
	if len(ali.Lines) == 0 {
		return m
	}
	for _, line := range beans {
		m[now] = line
		now = now.Add(line.Distance())
	}
	return m
}

// GetBeans 获取反向字典
// 比较顺序由字段顺序决定：
//  1. 先比较 Line.Time
//  2. 再比较 Line.A
//  3. 最后比较 Line.B
func (ali *AlibabaCompany) GetBeans(t time.Time, beans []model.Bean) map[model.Bean]time.Time {
	m := make(map[model.Bean]time.Time)
	now := t
	if len(ali.Lines) == 0 {
		return m
	}
	for _, line := range beans {
		m[line] = now
		now = now.Add(line.Distance())
	}
	return m
}

func (ali *AlibabaCompany) GetCost() time.Duration {
	d := time.Second * 0
	for _, line := range ali.Lines {
		d += line.Distance()
	}
	return d
}
