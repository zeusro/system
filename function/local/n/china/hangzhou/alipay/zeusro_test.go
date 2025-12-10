package alipay

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestCards(t *testing.T) {
	z := Zeusro{}
	// now := time.Now()
	// cards := z.Cards(now)
	// for _, c := range cards {
	// 	fmt.Println(c)
	// }
	deals := generateRandomDeals(10)
	payments := z.Pay(deals)
	for _, pay := range payments {
		fmt.Println(pay)
	}
}

// generateRandomDeals 返回随机交易记录切片，使用局部 rand.Rand 并加入额外随机因子
func generateRandomDeals(n int) []Deal {
	// 使用本地随机源，避免全局种子冲突并增加熵
	seed := time.Now().UnixNano() ^ int64(time.Now().Unix())
	rnd := rand.New(rand.NewSource(seed))

	deals := make([]Deal, n)

	for i := 0; i < n; i++ {
		// 随机生成过去 180~240 天内的交易（daysRange 带波动）
		daysRange := 180 + rnd.Intn(61) // 180..240
		daysAgo := rnd.Intn(daysRange)
		baseDate := time.Now().AddDate(0, 0, -daysAgo)

		// 随机时间（时:分:秒），以使时间更真实
		hour := rnd.Intn(24)
		min := rnd.Intn(60)
		sec := rnd.Intn(60)
		randomDate := time.Date(baseDate.Year(), baseDate.Month(), baseDate.Day(), hour, min, sec, 0, baseDate.Location())

		// 金额生成：基础均匀分布 + 小幅抖动 + 少量大额突发
		baseAmt := float32(rnd.Intn(2000-1)+1) + rnd.Float32() // 1 ~ 2000 + fractional
		jitter := 1 + float32(rnd.Float64()*0.2-0.1)           // ±10%
		amount := baseAmt * jitter

		// 小概率产生大额交易，模拟真实场景（如购物/付房租）
		if rnd.Intn(100) < 5 { // 5% 概率
			amount += float32(rnd.Intn(5000)) + 100
		}

		// 四舍五入到分
		amount = float32(math.Round(float64(amount)*100) / 100)

		deals[i] = Deal{
			t:     randomDate,
			Money: amount,
		}
	}

	return deals
}
