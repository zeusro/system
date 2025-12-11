package alipay

import (
	"fmt"
	"time"
)

type Card struct {
	Bank        string    //金融机构
	BillingDate time.Time //账单日
	LastDate    time.Time //最后还款日
}

func (c Card) String() string {
	return fmt.Sprintf("使用%s支付", c.Bank)
}

type Deal struct {
	t       time.Time
	Money   float32
	policy  string
	Payment Card
}

func (d Deal) String() string {
	return fmt.Sprintf("%v:[%s]%s￥%v", d.t, d.policy, d.Payment.String(), d.Money)
}

type Policy interface {
	Match(d Deal) bool
	Name() string
	MVP(d Deal) Policy
}
