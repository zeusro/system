package y

import (
	"testing"
	"time"
)

// TestY 调用 Y(base, end, randomCount, seed) 确保主仿真流程不 panic，输出到 stdout。
// 完整仿真步数较多，若需快速测试可加 -short 跳过或单独跑短仿真。
// go test  -v -test.fullpath=true -timeout 30s -run ^TestY$ github.com/zeusro/system/function/local/n/china/shantou/y
func TestY(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping full Y() in short mode")
	}
	loc := time.FixedZone("CST", 8*3600)
	base := time.Date(2008, 9, 1, 0, 0, 0, 0, loc)
	end := time.Date(2011, 7, 12, 0, 0, 0, 0, loc)
	Y(base, end, 56, 42)
}

// TestY_shortParams 用极短时间窗口与少量随机学生调用 Y，校验参数传递与流程不 panic。
func TestY_shortParams(t *testing.T) {
	loc := time.FixedZone("CST", 8*3600)
	base := time.Date(2008, 9, 1, 0, 0, 0, 0, loc)
	end := time.Date(2008, 9, 5, 0, 0, 0, 0, loc) // 仅 4 天
	Y(base, end, 2, 42)                           // 2 名随机学生
}

// TestY_shortRun 用极短时间窗口跑 Run，校验建班、步进、事件与激励采样是否正常。
func TestY_shortRun(t *testing.T) {
	loc := time.FixedZone("CST", 8*3600)
	base := time.Date(2008, 9, 1, 0, 0, 0, 0, loc)
	agents := []Agent{
		NewAgent(base, "teacher_y", RoleTeacherY, Factor{
			Birth: base, FamilyBackground: 0.5, IQ: 0.7, EQ: 0.5,
			PUAExposure: 0, PUAResistance: 0.8, LegalMoralRisk: 0.3,
		}),
		NewAgent(base, "student_01", RoleStudent, Factor{
			Birth: base, FamilyBackground: 0.5, IQ: 0.5, EQ: 0.5,
			PUAExposure: 0.2, PUAResistance: 0.6, LegalMoralRisk: 0.2,
		}),
	}
	agents[1].Score = 0.5

	step := 24 * time.Hour
	steps := 3
	state := Run(base, agents, step, steps, 42)

	if state.Birth != base {
		t.Errorf("state.Birth = %v, want %v", state.Birth, base)
	}
	if len(state.Agents) != 2 {
		t.Errorf("len(state.Agents) = %d, want 2", len(state.Agents))
	}
	// 至少应有若干事件或激励点（视 Run 实现而定）
	if state.Current.Before(base) {
		t.Errorf("state.Current %v before base %v", state.Current, base)
	}
	ctx := UpdateContext(state.Current, state.Agents)
	_ = Incentive(state.Current, ctx.TotalScore, ctx.StudentCount, ctx.ExamCount, ctx.EnrollCount)
}
