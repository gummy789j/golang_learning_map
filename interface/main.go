package main
// 此案例我希望能排序自訂義的structure利用interface
// 經典案例排列structure

import (
	"fmt"
	"sort"
	"math/rand"
)


	// Len is the number of elements in the collection.
	// Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
	
	// Note that floating-point comparison (the < operator on float32 or float64 values)
	// is not a transitive ordering when not-a-number (NaN) values are involved.
	// See Float64Slice.Less for a correct implementation for floating-point values.
	// Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	// Swap(i, j int)
type Hero struct {
	Name string
	Age int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

// 這邊我是想要以hero的age由小到大進行排序
// input i ,j return bool，幫助sort包判斷要從 i排到j 還是 j排到i
// 所以我為了幫助他判斷我的排序基準我告訴他 i位置hero的age 如果小於 j位置hero的age
// 若是true 那就代表 i位置hero的age 小於 j位置hero的age
// 那i要排在j前面 => 對我來說就可以得到由age小排到大的結果

func (hs HeroSlice) Less(i,j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i,j int) {
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	// 與上面是等價的
	hs[i] , hs[j] = hs[j] , hs[i]
}


func main() {

	var heros HeroSlice

	for i := 0; i < 10; i++ {
		hero := Hero{
			Name : fmt.Sprintf("Hero %d~", rand.Intn(100)),
			Age : rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	for _,v := range heros {
		fmt.Println(v)
	}
	fmt.Println("----------------------------")

	sort.Sort(heros)
	for _,v := range heros {
		fmt.Println(v)
	}
}


