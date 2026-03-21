package chenzhou

import "testing"

func TestPlateRegion(t *testing.T) {
	tests := []struct {
		plate    string
		wantProv string
		wantCity string
		wantOK   bool
	}{
		{"湘L88888", "湖南省", "郴州", true},
		{"粤B12345", "广东省", "深圳", true},
		{"京A12345", "北京市", "公交/城区", true},
		{" 沪 C·12345 ", "上海市", "郊区（禁入外环）", true},
		{"not-a-plate", "", "", false},
		{"湘12345", "", "", false},
		{"湘L653", "湖南省", "郴州", true},
	}
	for _, tt := range tests {
		prov, city, ok := PlateRegion(tt.plate)
		if ok != tt.wantOK || prov != tt.wantProv || city != tt.wantCity {
			t.Errorf("PlateRegion(%q) = (%q, %q, %v), want (%q, %q, %v)",
				tt.plate, prov, city, ok, tt.wantProv, tt.wantCity, tt.wantOK)
		}
	}
}
