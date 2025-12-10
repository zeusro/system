package alipay

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Zeusro struct {
}

func (z Zeusro) Pay(deals []Deal) []Deal {
	//follow the rule and find the best policy
	discount := NewDiscountPolicys(DefaultDiscountPolicys)
	for i, deal := range deals {
		if deal.Money > SmallMoney {
			//时间账单策略 BillingDatePolicy
			policy1 := NewBillingDatePolicy(z.Cards(deal.t))
			mvp := policy1.MVP(deal)
			best := mvp.BestCard
			deals[i].Payment = best
			deals[i].policy = policy1.Name()
		} else {
			//最大交易折扣交易策略 DiscountPolicys
			mvp := discount.MVP(deal)
			card := Card{Bank: mvp.Bank}
			deals[i].Payment = card
			deals[i].policy = mvp.Name()
		}
	}
	//TODO 账单日第二天消费策略 有bug
	return deals
}

// Cards 从本地文件加载信用卡信息，格式是每行 "银行 账单日/最后还款日"。如“农业银行 1/26”
func (z Zeusro) Cards(t time.Time) []Card {
	// 假设 card.txt 和 .go 文件在同一目录或可通过相对路径访问
	_, currentFile, _, _ := runtime.Caller(0)
	filename := filepath.Join(filepath.Dir(currentFile), "card.txt")
	f, err := os.Open(filename)
	if err != nil {
		panic("cannot open card.txt: " + err.Error())
	}
	defer f.Close()

	var cards []Card
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) != 2 {
			continue
		}
		bank := parts[0]
		days := parts[1]
		sl := strings.Split(days, "/")
		if len(sl) != 2 {
			continue
		}
		billingDay, err1 := strconv.Atoi(sl[0])
		lastDay, err2 := strconv.Atoi(sl[1])
		if err1 != nil || err2 != nil {
			continue
		}
		now := t
		billingDate := time.Date(now.Year(), now.Month(), billingDay, 0, 0, 0, 0, now.Location())
		lastDate := time.Date(now.Year(), now.Month(), lastDay, 0, 0, 0, 0, now.Location())
		//需要顺延一个月或者跨年
		if lastDay < billingDay {
			//2025/2026
			if now.Month() == time.December {
				lastDate = time.Date(now.Year()+1, time.January, lastDay, 0, 0, 0, 0, now.Location())
			} else {
				//  not 12/1
				lastDate = time.Date(now.Year(), now.Month()+1, lastDay, 0, 0, 0, 0, now.Location())
			}
		}
		cards = append(cards, Card{
			Bank:        bank,
			BillingDate: billingDate,
			LastDate:    lastDate,
		})
	}
	if err := scanner.Err(); err != nil {
		panic("error reading card.txt: " + err.Error())
	}
	return cards
}
