package y

import "time"

// Role 社会成员角色
type Role int

const (
	RoleTeacherY      Role = iota // 教师Y：以平均分激励，可减少分母
	RoleTeacherF                  // 教师F：平均分+本科升学率，可减少参考人数
	RoleStudentJudas              // 学生犹大：攀附教师F，网络暴力减员
	RoleStudentBlackMamba         // 学生黑曼巴：钞能力，可不参加高考（无策略，y.md）
	RoleStudentP                  // 学生P：低IQ，休学消极对抗
	RoleStudentY                  // 学生Y：高IQ运动员，加分需领导同意（y.md）
	RoleStudentC13                // 学生C13：高IQ贫困
	RoleStudent                   // 普通学生：随机样本，按因子选策略
	RolePsychologist               // 心理老师：降压力、安抚
	RoleSchoolLeader               // 学校领导：政绩=平均分+升学率，负责分配资源（如谁加分、谁休学退学）、安排心理老师定向辅导学生（y.md）
)

// NumStrategies 策略枚举数量，用于 StrategyCount 长度
const NumStrategies = 14

func (r Role) String() string {
	switch r {
	case RoleTeacherY:
		return "教师Y"
	case RoleTeacherF:
		return "教师F"
	case RoleStudentJudas:
		return "学生犹大"
	case RoleStudentBlackMamba:
		return "学生黑曼巴"
	case RoleStudentP:
		return "学生P"
	case RoleStudentY:
		return "学生Y"
	case RoleStudentC13:
		return "学生C13"
	case RoleStudent:
		return "普通学生"
	case RolePsychologist:
		return "心理老师"
	case RoleSchoolLeader:
		return "学校领导"
	default:
		return "未知"
	}
}

// Strategy 策略枚举：不同角色可采取的离散策略
type Strategy int

const (
	StrategyNone Strategy = iota
	// 教师
	StrategyPUA              // PUA施压减员
	StrategyReduceExamCount  // 减少高考参考人数
	StrategyLieEvade         // 撒谎与躲避监控
	StrategyNormalTeach      // 正常教学
	// 学生
	StrategyDropout          // 休学/退学
	StrategyNetworkViolence  // 网络暴力
	StrategyAthleteBonus     // 运动员加分
	StrategyStudyHard        // 努力学习
	StrategyAvoid            // 回避对抗
	// 心理
	StrategyDecompress       // 减压安抚
	// 领导
	StrategyPressureDown     // 向下施压（踢猫）
	StrategyIncentiveDesign   // 设计激励函数
)

func (s Strategy) String() string {
	switch s {
	case StrategyPUA:
		return "PUA施压减员"
	case StrategyReduceExamCount:
		return "减少高考参考人数"
	case StrategyLieEvade:
		return "撒谎躲避监控"
	case StrategyNormalTeach:
		return "正常教学"
	case StrategyDropout:
		return "休学退学"
	case StrategyNetworkViolence:
		return "网络暴力"
	case StrategyAthleteBonus:
		return "运动员加分"
	case StrategyStudyHard:
		return "努力学习"
	case StrategyAvoid:
		return "回避对抗"
	case StrategyDecompress:
		return "减压安抚"
	case StrategyPressureDown:
		return "向下施压"
	case StrategyIncentiveDesign:
		return "设计激励函数"
	default:
		return "无"
	}
}

// Agent 社会成员（时间序列对象：Birth 为第一成员）
type Agent struct {
	Birth time.Time // 时间第一成员：入校/入职时刻

	ID     string
	Role   Role
	Factor Factor // 量化因子

	// 状态（随时间演化）
	InSchool    bool    // 是否在校
	InExamPool  bool    // 是否在高考参考池
	Score       float64 // 当前成绩
	ScoreHistory [3]float64 // 过往3年年末成绩（0=最早年，2=最近年），用于高考成绩预测
	Stress      float64 // 心理压力 0~1
	LegalRisk   float64 // 当前行为导致的法规道德风险累积

	// 策略统计：仿真过程中各策略被选中的次数（仅学生有意义）
	StrategyCount []int

	// LastStrategy 上一期采取的策略（重复博弈：历史依赖）
	LastStrategy Strategy
}

// NewAgent 构造 Agent，时间必须参与初始化
func NewAgent(t time.Time, id string, role Role, f Factor) Agent {
	examPool := role != RoleStudentBlackMamba
	sc := make([]int, NumStrategies)
	return Agent{
		Birth:         t,
		ID:            id,
		Role:          role,
		Factor:        f,
		InSchool:      true,
		InExamPool:    examPool,
		Score:         0.5,
		Stress:        0,
		LegalRisk:     f.LegalMoralRisk,
		StrategyCount: sc,
	}
}
