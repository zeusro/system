package v1

import (
	"fmt"
	"sync"
	"testing"
)

func TestRichYoungAndBeautifulLookAtMe(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	var list1, list2, list3 []int64
	// 1 hour
	go func(w *sync.WaitGroup) {
		list1 = Zeusro(limit)
		wg.Done()
	}(&wg)
	go func(w *sync.WaitGroup) {
		list2 = YoungAndBeautiful(limit)
		wg.Done()
	}(&wg)
	go func(w *sync.WaitGroup) {
		list3 = RichGrandma(limit)
		wg.Done()
	}(&wg)
	wg.Wait()
	fmt.Printf("list1 %+v \t\n", list1)
	fmt.Printf("list2 %+v\t\n", list2)
	fmt.Printf("list3 %+v\t\n", list3)
	//TODO:没有LINQ 没有泛型，求个毛交集
	fmt.Println("MVP Equilibrium")
}
