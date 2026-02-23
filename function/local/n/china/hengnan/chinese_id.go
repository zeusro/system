package hengnan

import (
	"fmt"
	"strings"
)

// GB 11643-1999 身份证校验规则
// 加权因子表（从左到右，第1~17位）
var weight = [17]int{7, 9, 10, 5, 8, 4, 2, 6, 3, 7, 9, 10, 5, 8, 4, 2}

// 校验码对应表（模11的结果 → 校验码字符）
var verifyCodeTable = [11]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

func isValidChineseID(id string) bool {
	// 1. 长度必须是18位
	if len(id) != 18 {
		return false
	}

	// 2. 前17位必须是数字
	for i := 0; i < 17; i++ {
		if id[i] < '0' || id[i] > '9' {
			return false
		}
	}

	// 3. 第18位可以是 0-9 或 X/x
	lastChar := id[17]
	if !(lastChar >= '0' && lastChar <= '9') && lastChar != 'X' && lastChar != 'x' {
		return false
	}

	// 4. 计算加权和
	var sum int
	for i := 0; i < 17; i++ {
		digit := int(id[i] - '0')
		sum += digit * weight[i]
	}

	// 5. 模11
	mod := sum % 11

	// 6. 根据模的结果得到应该的校验码
	expected := verifyCodeTable[mod]

	// 7. 比较（X 和 x 都视为正确）
	actual := lastChar
	if actual == 'x' {
		actual = 'X'
	}

	return actual == expected
}

// 更友好的版本：返回校验是否通过 + 错误原因（可选）
func ValidateChineseID(id string) (bool, string) {
	id = strings.ToUpper(id) // 统一转为大写，方便处理 x/X

	if len(id) != 18 {
		return false, "身份证长度必须为18位"
	}

	// 检查前17位是否纯数字
	for i := 0; i < 17; i++ {
		if id[i] < '0' || id[i] > '9' {
			return false, "前17位必须全部为数字"
		}
	}

	// 最后一位只能是0-9或X
	if !(id[17] >= '0' && id[17] <= '9') && id[17] != 'X' {
		return false, "最后一位必须是0-9或X"
	}

	var sum int
	for i := 0; i < 17; i++ {
		sum += int(id[i]-'0') * weight[i]
	}

	mod := sum % 11
	expected := verifyCodeTable[mod]

	if id[17] != expected {
		return false, fmt.Sprintf("校验码错误，应为 %c，实际为 %c", expected, id[17])
	}

	return true, "校验通过"
}
