package v3

import (
	"fmt"
	"time"

	"github.com/zeusro/system/function/local/n/china/shenzhen/szx/model"
)

// Journey  <Sternstunden der Menschheit>
type Journey struct {
	Lines  []model.Line             //N维线段（为了简化运算不引入时间）的二维线段数组表示
	NBeans map[model.Bean]time.Time //N维对象
}

func NewJourney(n int) *Journey {
	//预分配内存减少消耗
	return &Journey{
		Lines:  make([]model.Line, n),
		NBeans: make(map[model.Bean]time.Time, n),
	}
}

func (j *Journey) AddLine(date time.Time, i int, line model.Line) {
	j.Lines[i] = line
	j.NBeans[model.Bean{Line: line}] = date
}

// Validate 以t1的B点和t2的A点进行连续性验证
func (j *Journey) Validate() (bool, error) {
	b := j.Lines[0].B
	for i := 1; i < len(j.Lines); i++ {
		a := j.Lines[i].A
		if a.Compare(b) {
			b = j.Lines[i].B
			continue
		}
		//最后一个点是(X.0,X.0) != (0.000000,0.000000)，直接跳过
		// fmt.Printf("dd%d\n", len(j.Lines))
		if i == len(j.Lines) {
			return true, nil
		}
		return false, fmt.Errorf("line %d 不连续：%v != %v", i, a, b)
	}
	return true, nil
}

// End
func (j *Journey) End() string {
	return "关于这趟旅程我还能说啥呢，总比宅在家里玩 Nintendo Switch 好多了"
}
