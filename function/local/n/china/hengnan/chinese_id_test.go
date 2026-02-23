package hengnan

import (
	"testing"
)

func TestValidateChineseID(t *testing.T) {
	tests := []struct {
		id     string
		wantOk bool
	}{
		{"430422199606039411", true},
		{"110101199003076219", false}, // 校验码错误
		{"32058119870616321X", false}, // 校验码错误
		{"440514198701084321", false}, // 故意错的校验码
		{"123456789012345678", false},
		{"43042219960603941", false},  // 长度不够(17位)
	}

	for _, tt := range tests {
		ok, _ := ValidateChineseID(tt.id)
		if ok != tt.wantOk {
			t.Errorf("ValidateChineseID(%q) = %v, want %v", tt.id, ok, tt.wantOk)
		}
	}
}
