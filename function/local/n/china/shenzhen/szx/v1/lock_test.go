package v1

import (
	"fmt"
	"testing"
)

func TestGetCost(t *testing.T) {
	lock := RWLock{}
	fmt.Printf("cost: %v", lock.GetCost())
}
