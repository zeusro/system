package v1

import (
	"fmt"

	"github.com/zeusro/system/function/local/n/china/shenzhen/szx/model"
)

// EatBeanWithLock 通过O(n)的读写锁解题
func EatBeanWithLock() {
	m := map[int]model.Point{}
	n := 50
	for i := 0; i < n; i++ {
		p := model.RandonPoint()
		m[i] = p
		// fmt.Println(p)
	}
	beans := NewBeans(m)
	//简化问题，把随机初始两点作为吃豆人起点
	a := make([]model.Point, 1)
	a[0] = m[0]
	beans.GetAndRemove(0)
	b := make([]model.Point, 1)
	b[0] = m[1]
	beans.GetAndRemove(1)
	alipay := AlibabaCompany{}
	aliyun := AlibabaCompany{}
	for i := 2; i < n; i++ {
		p, _ := beans.GetAndRemove(i)
		line1 := model.NewLine(a[len(a)-1], p)
		// fmt.Println(line1.Distance())
		line2 := model.NewLine(b[len(b)-1], p)
		// fmt.Println(line2.Distance())
		if line1.Distance() < line2.Distance() {
			// a = append(a, p)
			alipay.Lines = append(alipay.Lines, line1)
		} else {
			// b = append(b, p)
			aliyun.Lines = append(aliyun.Lines, line2)
		}
	}
	// fmt.Println(a)
	// fmt.Println(b)
	alipayBeans := make([]model.Bean, 0)
	for _, item := range alipay.Lines {
		alipayBeans = append(alipayBeans, model.Bean{Line: item})
	}
	beansA := alipay.EatBeans(alipayBeans)
	fmt.Println("支付宝吃豆人：")
	fmt.Println(beansA)

	aliyunBeans := make([]model.Bean, 0)
	for _, item := range aliyun.Lines {
		aliyunBeans = append(aliyunBeans, model.Bean{Line: item})
	}
	beansB := aliyun.EatBeans(aliyunBeans)
	fmt.Println("阿里云吃豆人：")
	fmt.Println(beansB)
}
