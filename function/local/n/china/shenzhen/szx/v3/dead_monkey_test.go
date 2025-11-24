package v3

import (
	"fmt"
	"testing"
	"time"
)

// go test -run TestFight -v
func TestFight(t *testing.T) {
	targets := []string{"a", "b", "c"}
	n := 50
	wukong := NewDeadMonkey(time.Now(), len(targets), n)
	wukong.Fight(targets)
	for k, v := range wukong.GoldenStaff {
		fmt.Println(v.String(k))
	}
	wukong.GoesToHell()
	if len(wukong.GoldenStaff) != (n + len(targets) - 2) {
		t.Logf("n:%v len:%v", n, len(wukong.GoldenStaff))
		t.FailNow()
	}
}
