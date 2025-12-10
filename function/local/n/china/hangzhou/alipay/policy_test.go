package alipay

import (
	"fmt"
	"testing"
	"time"
)

func TestBillingDatePolicy(t *testing.T) {
	now := time.Now()
	deal := Deal{
		t:     time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.Local),
		Money: 18,
	}
	b := BillingDatePolicy{
		cards: loadCards(now),
	}
	fmt.Println(b.MVP(deal).BestCard)
}

func TestDiscountPolicys(t *testing.T) {

	p2 := NewDiscountPolicys(DefaultDiscountPolicys)
	deal := Deal{
		t:     time.Now(),
		Money: 16,
	}
	p2.MVP(deal).String()

	deal = Deal{
		t:     time.Now(),
		Money: 18,
	}
	p2.MVP(deal).String()
	p2.MVP(deal).String()
}

func loadCards(t time.Time) []Card {
	z := Zeusro{}
	cards := z.Cards(t)
	return cards
}
