package problems

import (
	"fmt"
	"testing"
)

func TestGenerateUniqueSet(t *testing.T) {
	nums, target := GenerateSubsetSumProblem(10, 100)
	fmt.Println("输入集合:", nums)
	fmt.Println("目标和 T:", target)
}
