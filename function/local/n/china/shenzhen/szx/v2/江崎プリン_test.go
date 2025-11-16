package v2

import (
	"testing"

	"github.com/zeusro/system/function/local/n/china/shenzhen/szx/model"
)

func TestEatBean(t *testing.T) {
	ali := AlibabaGroup{}
	beans := make([]model.Bean, 100)
	ali.EatBean(beans)
}
