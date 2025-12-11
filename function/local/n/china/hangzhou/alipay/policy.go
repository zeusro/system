package alipay

import (
	"fmt"
	"time"
)

const SmallMoney float32 = 1000

var DefaultDiscountPolicys = []DiscountPolicy{
	{Money: 18, Discount: -0.3, Event: "每月返现（需要app每个月申请）", Bank: "农业银行", N: 1}, //todo 1天只能第一笔生效
	{Money: 16, Discount: -0.01, Event: "笔笔返现（需要app点击领取）", Bank: "浦发", N: -1},
	{Money: 1, Discount: -0.01, Event: "笔笔返现（需要app点击领取）", Bank: "广发银行", N: -1}, //todo，广发很抠，降低权重
}

// DiscountPolicy 基于单次小额交易折扣最大化交易策略
type DiscountPolicy struct {
	Money    float32 //交易门槛
	Discount float32 //折扣
	Event    string  //折扣名称
	Bank     string  //金融机构
	N        int     //-1表示n次，大于0表示有限次数
}

func (r DiscountPolicy) Match(d Deal) bool {
	if r.N == 0 {
		return false
	}
	if d.Money > SmallMoney || d.Money < r.Money {
		return false
	}
	return true
}

func (p DiscountPolicy) Name() string {
	return "小额优惠交易"
}

func (p DiscountPolicy) String() {
	fmt.Printf("%s-%s: 满%.2f减%.2f\n", p.Bank, p.Name(), p.Money, p.Discount)
}

type DiscountPolicys struct {
	policys  []DiscountPolicy
	Resource map[string]DiscountPolicy
}

func NewDiscountPolicys(policys []DiscountPolicy) DiscountPolicys {
	rules := DiscountPolicys{policys: policys}
	rules.Resource = rules.ToDiscountPolicyMap()
	return rules
}

func (p DiscountPolicys) ToDiscountPolicyMap() map[string]DiscountPolicy {
	m := make(map[string]DiscountPolicy)
	for _, r := range p.policys {
		m[r.Bank] = r
	}
	return m
}

func (p *DiscountPolicys) MVP(d Deal) DiscountPolicy {
	var best DiscountPolicy
	var maxDiscount float32 = 0
	for _, policy := range p.Resource {
		if !policy.Match(d) {
			continue
		}
		if policy.Discount >= maxDiscount {
			continue
		}
		best = policy
		if best.N > 0 {
			maxDiscount = best.Discount
			//资源池减一
			best.N--
		}
	}
	if len(best.Bank) > 0 {
		p.Resource[best.Bank] = best
	}
	return best
}

// BillingDatePolicy “第二”账单日消费策略
type BillingDatePolicy struct {
	cards    []Card
	BestCard Card
}

func NewBillingDatePolicy(cards []Card) BillingDatePolicy {
	p := BillingDatePolicy{cards: cards}
	return p
}

func (p BillingDatePolicy) Match(d Deal) bool {
	return true
}

func (p BillingDatePolicy) MVP(d Deal) BillingDatePolicy {
	if len(p.cards) == 0 {
		return p
	}

	now := d.t
	var bestCard *Card
	var bestDelta time.Duration = time.Hour * 24 * 31 // 极大值

	for i := range p.cards {
		c := &p.cards[i]
		delta := now.Sub(c.BillingDate.AddDate(0, 0, 1))
		if delta > 0 && delta < bestDelta {
			bestDelta = delta
			bestCard = c
		}
	}
	if bestCard != nil {
		p.BestCard = *bestCard
	}
	return p
}

func (p BillingDatePolicy) Name() string {
	return "账单日第二天消费策略"
}
