package guangzhou

import "testing"

// --- PASS: TestCoke (0.00s)
// PASS
// ok  	github.com/zeusro/system/function/local/n/china/guangzhou	0.330s
func TestCoke(t *testing.T) {
	// 附近商户卖[3.39 3.43 3.38 3.34 3.58 3.48 3.56 3.4 3.5 3.31 3.3 3.53 3.51 3.33]
	// 卷王卖3.3,我卖4.08
	Coke(2.5, 4, 14, 0.5)
}
